package logic

import (
	"cloud-disk/common/errorx"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest) (resp *types.UserFileMoveResponse, err error) {
	//检测该文件是否存在
	userFileInfo, err := l.svcCtx.UserRepositoryModel.FindByIdentity(l.ctx, req.Identity)
	if err != nil {
		return nil, errorx.NewDefaultError("原文件不存在！")
	}
	//检测新目录是否已存在该文件
	count, err := l.svcCtx.UserRepositoryModel.CountByIdentityAndParentId(l.ctx, req.Identity, json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String(), req.ParentId)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewDefaultError("已存在相同名称的文件！")
	}
	//修改
	userFileInfo.ParentId = sql.NullInt64{
		Int64: req.ParentId,
		Valid: true,
	}
	err = l.svcCtx.UserRepositoryModel.Update(l.ctx, userFileInfo)
	if err != nil {
		return nil, err
	}
	return &types.UserFileMoveResponse{}, nil
}
