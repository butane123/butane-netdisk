package logic

import (
	"butane-netdisk/service/repository/rpc/internal/svc"
	"butane-netdisk/service/repository/rpc/types/repository"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepositoryPoolByRepositoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepositoryPoolByRepositoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepositoryPoolByRepositoryIdLogic {
	return &GetRepositoryPoolByRepositoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepositoryPoolByRepositoryIdLogic) GetRepositoryPoolByRepositoryId(in *repository.RepositoryReq) (*repository.RepositoryResp, error) {
	repositoryPoolInfo, err := l.svcCtx.RepositoryPoolModel.FindOne(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &repository.RepositoryResp{
		Ext:  repositoryPoolInfo.Ext,
		Size: repositoryPoolInfo.Size,
		Path: repositoryPoolInfo.Path,
		Name: repositoryPoolInfo.Name,
	}, nil
}
