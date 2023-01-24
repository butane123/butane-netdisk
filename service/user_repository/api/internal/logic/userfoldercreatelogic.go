package logic

import (
	"cloud-disk/common/utils"
	"cloud-disk/service/user_repository/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"

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
	existCount, err := l.svcCtx.UserRepositoryModel.CountByParentIdAndName(l.ctx, req.ParentId, json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String(), req.Name)
	if err != nil {
		return nil, errors.New("验证文件夹名字不存在失败！")
	}
	if existCount > 0 {
		return nil, errors.New("已存在相同名称的文件夹！")
	}
	newIdentity := utils.GenerateUUID()
	_, err = l.svcCtx.UserRepositoryModel.Insert(l.ctx, &model.UserRepository{
		Identity:     sql.NullString{newIdentity, true},
		UserIdentity: sql.NullString{json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String(), true},
		ParentId:     sql.NullInt64{req.ParentId, true},
		Name:         sql.NullString{req.Name, true},
	})
	if err != nil {
		return nil, err
	}
	return &types.UserFolderCreateResponse{Identity: newIdentity}, nil
}
