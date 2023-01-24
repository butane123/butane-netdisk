package model

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepositoryPoolModel = (*customRepositoryPoolModel)(nil)

type (
	// RepositoryPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepositoryPoolModel.
	RepositoryPoolModel interface {
		repositoryPoolModel
		FindRepositoryPoolByRepositoryId(ctx context.Context, repositoryId string) (*RepositoryPool, error)
		DeleteByIdentity(ctx context.Context, repositoryId string) (*RepositoryPool, error)
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

func (m *defaultRepositoryPoolModel) FindRepositoryPoolByRepositoryId(ctx context.Context, repositoryId string) (*RepositoryPool, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("identity = ?", repositoryId).Where("deleted_at is null").ToSql()
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

func (m *defaultRepositoryPoolModel) DeleteByIdentity(ctx context.Context, repositoryId string) (*RepositoryPool, error) {
	RepositoryInfo, err := m.FindRepositoryPoolByRepositoryId(ctx, repositoryId)
	switch err {
	case nil:
		break
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
	err = m.Delete(ctx, RepositoryInfo.Id)
	switch err {
	case nil:
		return RepositoryInfo, err
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRepositoryPoolModel) CountByHash(ctx context.Context, hash string) (int64, error) {
	countBuilder := m.CountBuilder("id")
	query, values, err := countBuilder.Where("hash = ?", hash).Where("deleted_at is null").ToSql()
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
	query, values, err := rowBuilder.Where("hash = ?", hash).Where("deleted_at is null").ToSql()
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
