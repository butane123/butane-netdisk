package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"butane-netdisk/service/user/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	fmt.Println("dadadadad")
	//1 验证输入的验证码是否正确
	verificationCode, err := l.svcCtx.RedisClient.Get(utils.CacheEmailCodeKey + req.Email)
	if err != nil || verificationCode == "" {
		return nil, errorx.NewDefaultError("无发送验证码或验证码已到期！")
	}
	if verificationCode != req.Code {
		return nil, errorx.NewDefaultError("输入的验证码不一致！")
	}
	//2 检查用户名是否重复
	_, err = l.svcCtx.UserBasicModel.FindByName(l.ctx, req.Name)
	switch err {
	case nil:
		return nil, errorx.NewDefaultError("用户名重复，请重试！")
	case sqlc.ErrNotFound:
		break
	default:
		return nil, err
	}
	//3 生成新用户，插入数据库中
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "user")
	_, err = l.svcCtx.UserBasicModel.InsertWithId(l.ctx, &model.UserBasic{
		Id:       newId,
		Name:     req.Name,
		Password: utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},
	},
	)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{}, nil
}
