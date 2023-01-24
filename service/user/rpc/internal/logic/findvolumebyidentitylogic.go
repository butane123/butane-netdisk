package logic

import (
	"cloud-disk/service/user/rpc/internal/svc"
	"cloud-disk/service/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVolumeByIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVolumeByIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVolumeByIdentityLogic {
	return &FindVolumeByIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVolumeByIdentityLogic) FindVolumeByIdentity(in *user.FindVolumeReq) (*user.FindVolumeResp, error) {
	userInfo, err := l.svcCtx.UserBasicModel.FindByIdentity(l.ctx, in.Identity)
	if err != nil {
		return nil, err
	}
	return &user.FindVolumeResp{
		NowVolume:   userInfo.NowVolume,
		TotalVolume: userInfo.TotalVolume,
	}, nil
}
