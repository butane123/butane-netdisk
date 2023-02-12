package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"butane-netdisk/service/repository/rpc/types/repository"
	"butane-netdisk/service/user/rpc/types/user"
	"butane-netdisk/service/user_repository/model"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest) (resp *types.UserRepositorySaveResponse, err error) {
	// 判断文件是否超容量
	// //获取文件Size
	repositoryPoolInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: req.RepositoryId})
	if err != nil {
		return nil, err
	}
	// //更新个人total_Volume、nowVolume
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	userInfo, err := l.svcCtx.UserRpc.FindVolumeById(l.ctx, &user.FindVolumeReq{Id: userId})
	if err != nil {
		return nil, err
	}
	if repositoryPoolInfo.Size+userInfo.NowVolume > userInfo.TotalVolume {
		return nil, errorx.NewDefaultError("文件超出容量限制！")
	}
	// 更新当前容量
	_, err = l.svcCtx.UserRpc.AddVolume(l.ctx, &user.AddVolumeReq{
		Id:   userId,
		Size: repositoryPoolInfo.Size,
	})
	if err != nil {
		return nil, err
	}
	// 新增关联记录
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "user_repository")
	_, err = l.svcCtx.UserRepositoryModel.Insert(l.ctx, &model.UserRepository{
		Id:           newId,
		UserId:       userId,
		ParentId:     req.ParentId,
		RepositoryId: req.RepositoryId,
		Name:         req.Name,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("存储失败！")
	}
	return &types.UserRepositorySaveResponse{}, nil
}
