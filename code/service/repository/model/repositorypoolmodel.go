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

var _ RepositoryPoolModel = (*customRepositoryPoolModel)(nil)

var repositoryPoolRowsExpectAutoSetButId = strings.Join(stringx.Remove(repositoryPoolFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")

type (
	// RepositoryPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepositoryPoolModel.
	RepositoryPoolModel interface {
		repositoryPoolModel
		InsertWithId(ctx context.Context, data *RepositoryPool) (sql.Result, error)
		CountByHash(ctx context.Context, hash string) (int64, error)
		FindRepositoryPoolByHash(ctx context.Context, hash string) (*RepositoryPool, error)
	}

	customRepositoryPoolModel struct {
		*defaultRepositoryPoolModel
	}
)

// NewRepositoryPoolModel returns a model for the database table.
func NewRepositoryPoolModel(conn sqlx.SqlConn, c cache.CacheConf) RepositoryPoolModel {
	return &customRepositoryPoolModel{
		defaultRepositoryPoolModel: newRepositoryPoolModel(conn, c),
	}
}

func (m *defaultRepositoryPoolModel) InsertWithId(ctx context.Context, data *RepositoryPool) (sql.Result, error) {
	repositoryPoolHashKey := fmt.Sprintf("%s%v", cacheRepositoryPoolHashPrefix, data.Hash)
	repositoryPoolIdKey := fmt.Sprintf("%s%v", cacheRepositoryPoolIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, repositoryPoolRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Hash, data.Ext, data.Size, data.Path, data.Name)
	}, repositoryPoolHashKey, repositoryPoolIdKey)
	return ret, err
}

func (m *defaultRepositoryPoolModel) CountByHash(ctx context.Context, hash string) (int64, error) {
	countBuilder := m.CountBuilder("id")
	query, values, err := countBuilder.Where("hash = ?", hash).ToSql()
	if err != nil {
		return 0, err
	}
	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultRepositoryPoolModel) FindRepositoryPoolByHash(ctx context.Context, hash string) (*RepositoryPool, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("hash = ?", hash).ToSql()
	if err != nil {
		return nil, err
	}
	var resp RepositoryPool
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

func (m *defaultRepositoryPoolModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(repositoryPoolRows).From(m.table)
}

func (m *defaultRepositoryPoolModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
