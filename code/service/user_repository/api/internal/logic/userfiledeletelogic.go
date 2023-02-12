package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/service/repository/rpc/types/repository"
	"butane-netdisk/service/user/rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest) (resp *types.UserFileDeleteResponse, err error) {
	//先删user_repository
	userFileInfo, err := l.svcCtx.UserRepositoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserRepositoryModel.Delete(l.ctx, userFileInfo.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("更新个人存储池失败！")
	}
	//从中心存储池中取size
	repositoryInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: userFileInfo.RepositoryId})
	if err != nil {
		return nil, errorx.NewDefaultError("中心存储池找不到该数据！")
	}
	//更新user_basic的now_volume
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.DecreaseVolume(l.ctx, &user.DecreaseVolumeReq{
		Id:   userId,
		Size: repositoryInfo.Size,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新容量失败！")
	}
	return &types.UserFileDeleteResponse{}, nil
}
