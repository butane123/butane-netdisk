package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"butane-netdisk/service/user_repository/model"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest) (resp *types.UserFolderCreateResponse, err error) {
	//验证文件夹名字不存在：
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	existCount, err := l.svcCtx.UserRepositoryModel.CountByParentIdAndName(l.ctx, req.ParentId, userId, req.Name)
	if err != nil {
		return nil, errorx.NewDefaultError("验证文件夹名字不存在失败！")
	}
	if existCount > 0 {
		return nil, errorx.NewDefaultError("已存在相同名称的文件夹！")
	}
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "user_repository")
	_, err = l.svcCtx.UserRepositoryModel.Insert(l.ctx, &model.UserRepository{
		Id:       newId,
		UserId:   userId,
		ParentId: req.ParentId,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserFolderCreateResponse{Id: newId}, nil
}
