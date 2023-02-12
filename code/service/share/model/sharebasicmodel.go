package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShareBasicModel = (*customShareBasicModel)(nil)

var shareBasicRowsExpectAutoSetButId = strings.Join(stringx.Remove(shareBasicFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// ShareBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShareBasicModel.
	ShareBasicModel interface {
		shareBasicModel
		InsertWithId(ctx context.Context, data *ShareBasic) (sql.Result, error)
		AddOneClick(ctx context.Context, id int64) error
	}

	customShareBasicModel struct {
		*defaultShareBasicModel
	}
)

// NewShareBasicModel returns a model for the database table.
func NewShareBasicModel(conn sqlx.SqlConn, c cache.CacheConf) ShareBasicModel {
	return &customShareBasicModel{
		defaultShareBasicModel: newShareBasicModel(conn, c),
	}
}
func (m *defaultShareBasicModel) InsertWithId(ctx context.Context, data *ShareBasic) (sql.Result, error) {
	shareBasicIdKey := fmt.Sprintf("%s%v", cacheShareBasicIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, shareBasicRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.RepositoryId, data.UserRepositoryId, data.ExpiredTime, data.ClickNum)
	}, shareBasicIdKey)
	return ret, err
}
func (m *defaultShareBasicModel) AddOneClick(ctx context.Context, id int64) error {
	shareBasicIdKey := fmt.Sprintf("%s%v", cacheShareBasicIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set click_num = click_num + 1 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, shareBasicIdKey)
	return err
}

func (m *defaultShareBasicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(shareBasicRows).From(m.table)
}
