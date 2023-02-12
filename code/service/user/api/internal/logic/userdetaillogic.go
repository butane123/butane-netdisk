package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	userInfo, err := l.svcCtx.UserBasicModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("id输入有误！")
	}
	return &types.DetailResponse{
		Name:  userInfo.Name,
		Email: userInfo.Email.String,
	}, nil
}
