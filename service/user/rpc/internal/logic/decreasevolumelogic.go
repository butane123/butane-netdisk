package logic

import (
	"context"

	"cloud-disk/service/user/rpc/internal/svc"
	"cloud-disk/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecreaseVolumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecreaseVolumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecreaseVolumeLogic {
	return &DecreaseVolumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecreaseVolumeLogic) DecreaseVolume(in *user.DecreaseVolumeReq) (*user.DecreaseVolumeResp, error) {
	userInfo, err := l.svcCtx.UserBasicModel.FindByIdentity(l.ctx, in.Identity)
	if err != nil {
		return nil, err
	}
	userInfo.NowVolume = userInfo.NowVolume - in.Size
	err = l.svcCtx.UserBasicModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &user.DecreaseVolumeResp{}, nil
}
