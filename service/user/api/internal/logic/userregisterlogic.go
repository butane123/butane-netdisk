package logic

import (
	"cloud-disk/common/errorx"
	"cloud-disk/common/utils"
	"cloud-disk/service/user/model"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"cloud-disk/service/user/api/internal/svc"
	"cloud-disk/service/user/api/internal/types"

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
	//1 验证输入的验证码是否正确
	verificationCode, err := l.svcCtx.RedisClient.Get(req.Email)
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
	_, err = l.svcCtx.UserBasicModel.Insert(l.ctx, &model.UserBasic{
		Identity: sql.NullString{String: utils.GenerateUUID(), Valid: true},
		Name:     sql.NullString{String: req.Name, Valid: true},
		Password: sql.NullString{String: req.Password, Valid: true},
		Email:    sql.NullString{String: req.Email, Valid: true},
	},
	)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{}, nil
}
