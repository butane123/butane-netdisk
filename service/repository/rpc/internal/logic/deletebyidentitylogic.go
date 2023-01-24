package logic

import (
	"context"

	"cloud-disk/service/repository/rpc/internal/svc"
	"cloud-disk/service/repository/rpc/repository"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteByIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteByIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteByIdentityLogic {
	return &DeleteByIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteByIdentityLogic) DeleteByIdentity(in *repository.DeleteByIdentityReq) (*repository.DeleteByIdentityResp, error) {
	repositoryInfo, err := l.svcCtx.RepositoryPoolModel.DeleteByIdentity(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &repository.DeleteByIdentityResp{Size: repositoryInfo.Size.Int64}, nil
}
