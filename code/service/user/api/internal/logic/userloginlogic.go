package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"context"
	"strings"
	"time"

	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1、从数据库中查询当前用户
	name, password := req.Name, req.Password
	if len(strings.TrimSpace(name)) == 0 || len(strings.TrimSpace(password)) == 0 {
		return nil, errorx.NewDefaultError("用户名或密码不能为空！")
	}
	userInfo, err := l.svcCtx.UserBasicModel.JudgeUserExist(l.ctx, name, utils.PasswordEncrypt(l.svcCtx.Config.Salt, password))
	if err != nil {
		return nil, errorx.NewDefaultError("无此用户或用户名与密码不匹配！")
	}
	// 2、生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccesssToken: jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
