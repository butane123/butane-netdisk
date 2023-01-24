package logic

import (
	"cloud-disk/service/user/rpc/internal/svc"
	"cloud-disk/service/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVolumeLogic {
	return &AddVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVolumeLogic) AddVolume(in *user.AddVolumeReq) (*user.AddVolumeResp, error) {
	userInfo, err := l.svcCtx.UserBasicModel.FindByIdentity(l.ctx, in.Identity)
	if err != nil {
		return nil, err
	}
	userInfo.NowVolume = userInfo.NowVolume + in.Size
	err = l.svcCtx.UserBasicModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &user.AddVolumeResp{}, nil
}
