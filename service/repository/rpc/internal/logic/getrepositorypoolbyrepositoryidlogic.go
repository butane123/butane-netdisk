package logic

import (
	"cloud-disk/service/repository/rpc/internal/svc"
	"cloud-disk/service/repository/rpc/repository"
	"context"
	"fmt"

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
	fmt.Println(in.RepositoryId)
	repositoryPoolInfo, err := l.svcCtx.RepositoryPoolModel.FindRepositoryPoolByRepositoryId(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &repository.RepositoryResp{
		Ext:  repositoryPoolInfo.Ext.String,
		Size: repositoryPoolInfo.Size.Int64,
		Path: repositoryPoolInfo.Path.String,
	}, nil
}
