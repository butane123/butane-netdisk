package logic

import (
	"context"

	"github.com/pkg/errors"

	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"

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
	//根据文件夹identity，获取id，然后作为父目录id去搜目录下的数据
	userRepositoryInfo, err := l.svcCtx.UserRepositoryModel.FindByIdentity(l.ctx, req.Identity)
	if err != nil {
		return nil, errors.New("文件夹identity有误！")
	}
	allUserRepository, err := l.svcCtx.UserRepositoryModel.FindAllFolderByParentId(l.ctx, userRepositoryInfo.Id, userRepositoryInfo.UserIdentity.String)
	if err != nil {
		return nil, errors.New("该文件夹下搜索文件夹失败！")
	}
	newList := make([]*types.UserFolder, 0)
	for _, userRepository := range allUserRepository {
		newList = append(newList, &types.UserFolder{
			Identity: userRepository.Identity.String,
			Name:     userRepository.Name.String,
		})
	}
	return &types.UserFolderListResponse{List: newList}, nil
}
