package logic

import (
	"cloud-disk/common/errorx"
	"cloud-disk/service/user/api/internal/svc"
	"cloud-disk/service/user/api/internal/types"
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
	userInfo, err := l.svcCtx.UserBasicModel.FindByIdentity(l.ctx, req.Identity)
	if err != nil {
		return nil, errorx.NewDefaultError("Identity输入有误！")
	}
	return &types.DetailResponse{
		Name:  userInfo.Name.String,
		Email: userInfo.Email.String,
	}, nil
}
