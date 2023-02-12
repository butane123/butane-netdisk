package logic

import (
	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest) (resp *types.UserFileNameUpdateResponse, err error) {
	userFileInfo, err := l.svcCtx.UserRepositoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	userFileInfo.Name = req.Name
	err = l.svcCtx.UserRepositoryModel.Update(l.ctx, userFileInfo)
	if err != nil {
		return nil, err
	}
	return &types.UserFileNameUpdateResponse{}, nil
}
