package logic

import (
	"context"

	"butane-netdisk/service/user/rpc/internal/svc"
	"butane-netdisk/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVolumeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVolumeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVolumeByIdLogic {
	return &FindVolumeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVolumeByIdLogic) FindVolumeById(in *user.FindVolumeReq) (*user.FindVolumeResp, error) {
	userInfo, err := l.svcCtx.UserBasicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.FindVolumeResp{
		NowVolume:   userInfo.NowVolume,
		TotalVolume: userInfo.TotalVolume,
	}, nil
}
