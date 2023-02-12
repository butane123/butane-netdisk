package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"

	"github.com/Masterminds/squirrel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserBasicModel = (*customUserBasicModel)(nil)

var userBasicRowsExpectAutoSetButIdAndTotalVolume = strings.Join(stringx.Remove(userBasicFieldNames, "`total_volume`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")

type (
	// UserBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBasicModel.
	UserBasicModel interface {
		userBasicModel
		InsertWithId(ctx context.Context, data *UserBasic) (sql.Result, error)
		JudgeUserExist(ctx context.Context, name string, password string) (*UserBasic, error)
		FindByName(ctx context.Context, name string) (*UserBasic, error)
		UpdateVolume(ctx context.Context, id int64, size int64) (result sql.Result, err error)
	}

	customUserBasicModel struct {
		*defaultUserBasicModel
	}
)

// NewUserBasicModel returns a model for the database table.
func NewUserBasicModel(conn sqlx.SqlConn, c cache.CacheConf) UserBasicModel {
	return &customUserBasicModel{
		defaultUserBasicModel: newUserBasicModel(conn, c),
	}
}

func (m *defaultUserBasicModel) InsertWithId(ctx context.Context, data *UserBasic) (sql.Result, error) {
	userBasicEmailKey := fmt.Sprintf("%s%v", cacheUserBasicEmailPrefix, data.Email)
	userBasicIdKey := fmt.Sprintf("%s%v", cacheUserBasicIdPrefix, data.Id)
	userBasicNameKey := fmt.Sprintf("%s%v", cacheUserBasicNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userBasicRowsExpectAutoSetButIdAndTotalVolume)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.Password, data.Email, data.NowVolume)
	}, userBasicEmailKey, userBasicIdKey, userBasicNameKey)
	return ret, err
}

func (m *defaultUserBasicModel) JudgeUserExist(ctx context.Context, name string, password string) (*UserBasic, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("name = ?", name).Where("password = ?", password).ToSql()
	var resp UserBasic
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserBasicModel) FindByName(ctx context.Context, name string) (*UserBasic, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("name = ?", name).ToSql()
	var resp UserBasic
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserBasicModel) UpdateVolume(ctx context.Context, id int64, size int64) (result sql.Result, err error) {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	userBasicEmailKey := fmt.Sprintf("%s%v", cacheUserBasicEmailPrefix, data.Email)
	userBasicIdKey := fmt.Sprintf("%s%v", cacheUserBasicIdPrefix, data.Id)
	userBasicNameKey := fmt.Sprintf("%s%v", cacheUserBasicNamePrefix, data.Name)
	res, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set now_volume = now_volume + ? where `id` = ? and `now_volume` + ? <= `total_volume`", m.table)
		return conn.ExecCtx(ctx, query, size, id, size)
	}, userBasicEmailKey, userBasicIdKey, userBasicNameKey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *defaultUserBasicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userBasicRows).From(m.table)
}

func (m *defaultUserBasicModel) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(m.table)
}
