package model

import (
	"cloud-disk/service/user/model"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRepositoryModel = (*customUserRepositoryModel)(nil)

type (
	// UserRepositoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRepositoryModel.
	UserRepositoryModel interface {
		userRepositoryModel
		FindByRepositoryId(ctx context.Context, repositoryId string) (*UserRepository, error)
		FindByIdentity(ctx context.Context, identity string) (*UserRepository, error)
		FindAllInPage(ctx context.Context, parentId int64, userIdentity string, startIndex int64, pageSize int64) ([]*UserRepository, error)
		FindAllFolderByParentId(ctx context.Context, parentId int64, userIdentity string) ([]*UserRepository, error)
		CountByParentIdAndName(ctx context.Context, parentId int64, userIdentity string, Name string) (int64, error)
		CountByIdentityAndParentId(ctx context.Context, identity string, userIdentity string, parentId int64) (int64, error)
	}

	customUserRepositoryModel struct {
		*defaultUserRepositoryModel
	}
)

// NewUserRepositoryModel returns a model for the database table.
func NewUserRepositoryModel(conn sqlx.SqlConn, c cache.CacheConf) UserRepositoryModel {
	return &customUserRepositoryModel{
		defaultUserRepositoryModel: newUserRepositoryModel(conn, c),
	}
}

func (m *defaultUserRepositoryModel) FindByRepositoryId(ctx context.Context, repositoryId string) (*UserRepository, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("repository_identity = ?", repositoryId).Where("deleted_at is null").ToSql()
	var resp UserRepository
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRepositoryModel) FindByIdentity(ctx context.Context, identity string) (*UserRepository, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("identity = ?", identity).Where("deleted_at is null").ToSql()
	var resp UserRepository
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserRepositoryModel) FindAllInPage(ctx context.Context, parentId int64, userIdentity string, startIndex int64, pageSize int64) ([]*UserRepository, error) {
	var resp []*UserRepository
	var resp1 UserRepository
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("parent_id = ?", fmt.Sprintf("%d", parentId)).Where("user_identity = ?", userIdentity).Where("deleted_at is null").Offset(uint64(startIndex)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowNoCacheCtx(ctx, &resp1, query, values...)
	if err != nil {
		return nil, err
	}
	//var respId []*int64
	//Sql := "select id from " + m.table
	//err = m.QueryRowNoCacheCtx(ctx, &respId, Sql)
	//fmt.Println(respId)  //返回[]，故无法多行查询
	//而gorm却可以
	resp = append(resp, &resp1)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRepositoryModel) FindAllFolderByParentId(ctx context.Context, parentId int64, userIdentity string) ([]*UserRepository, error) {
	var resp []*UserRepository
	var resp1 UserRepository
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("parent_id = ?", fmt.Sprintf("%d", parentId)).Where("user_identity = ?", userIdentity).Where("deleted_at is null").ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowNoCacheCtx(ctx, &resp1, query, values...)
	resp = append(resp, &resp1)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserRepositoryModel) CountByParentIdAndName(ctx context.Context, parentId int64, userIdentity string, Name string) (int64, error) {
	countBuilder := m.CountBuilder("id")
	query, values, err := countBuilder.Where("parent_id = ?", fmt.Sprintf("%d", parentId)).Where("user_identity = ?", userIdentity).Where("name = ?", Name).Where("deleted_at is null").ToSql()
	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, model.ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserRepositoryModel) CountByIdentityAndParentId(ctx context.Context, identity string, userIdentity string, parentId int64) (int64, error) {
	countBuilder := m.CountBuilder("id")
	query, values, err := countBuilder.Where("parent_id = ?", fmt.Sprintf("%d", parentId)).Where("user_identity = ?", userIdentity).Where("identity = ?", identity).Where("deleted_at is null").ToSql()
	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, model.ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserRepositoryModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRepositoryRows).From(m.table)
}

func (m *defaultUserRepositoryModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
