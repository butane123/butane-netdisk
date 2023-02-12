package logic

import (
	"butane-netdisk/common/errorx"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderListLogic {
	return &UserFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderListLogic) UserFolderList(req *types.UserFolderListRequest) (resp *types.UserFolderListResponse, err error) {
	//根据文件夹id，然后作为父目录id去搜目录下的数据
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	allUserRepository, err := l.svcCtx.UserRepositoryModel.FindAllFolderByParentId(l.ctx, req.Id, userId)
	if err != nil {
		return nil, errorx.NewDefaultError("该文件夹下搜索文件夹失败！")
	}
	newList := make([]*types.UserFolder, 0)
	for _, userRepository := range allUserRepository {
		newList = append(newList, &types.UserFolder{
			Id:   userRepository.Id,
			Name: userRepository.Name,
		})
	}
	return &types.UserFolderListResponse{List: newList}, nil
}
