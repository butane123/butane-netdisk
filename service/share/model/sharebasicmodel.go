package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShareBasicModel = (*customShareBasicModel)(nil)

type (
	// ShareBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShareBasicModel.
	ShareBasicModel interface {
		shareBasicModel
		AddOneClick(ctx context.Context, identity string) (*ShareBasic, error)
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

func (m *defaultShareBasicModel) AddOneClick(ctx context.Context, identity string) (*ShareBasic, error) {
	shareBasicInfo, err := m.FindByIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}
	shareBasicInfo.ClickNum++
	err = m.Update(ctx, shareBasicInfo)
	if err != nil {
		return nil, err
	}
	return shareBasicInfo, err
}

func (m *defaultShareBasicModel) FindByIdentity(ctx context.Context, identity string) (*ShareBasic, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("identity = ?", identity).ToSql()
	var resp ShareBasic
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

func (m *defaultShareBasicModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(shareBasicRows).From(m.table)
}
