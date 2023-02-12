package logic

import (
	"context"

	"butane-netdisk/service/repository/rpc/internal/svc"
	"butane-netdisk/service/repository/rpc/types/repository"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteByIdLogic {
	return &DeleteByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteByIdLogic) DeleteById(in *repository.DeleteByIdReq) (*repository.DeleteByIdResp, error) {
	repositoryInfo, err := l.svcCtx.RepositoryPoolModel.FindOne(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RepositoryPoolModel.Delete(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &repository.DeleteByIdResp{Size: repositoryInfo.Size}, nil
}
