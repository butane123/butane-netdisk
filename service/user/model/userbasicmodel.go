package model

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserBasicModel = (*customUserBasicModel)(nil)

type (
	// UserBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBasicModel.
	UserBasicModel interface {
		userBasicModel
		JudgeUserExist(ctx context.Context, name string, password string) (*UserBasic, error)
		FindByIdentity(ctx context.Context, identity string) (*UserBasic, error)
		FindByName(ctx context.Context, name string) (*UserBasic, error)
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

func (m *defaultUserBasicModel) FindByIdentity(ctx context.Context, identity string) (*UserBasic, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("identity = ?", identity).ToSql()
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

func (m *defaultUserBasicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userBasicRows).From(m.table)
}
