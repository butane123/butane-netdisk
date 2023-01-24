package logic

import (
	"cloud-disk/common/utils"
	"cloud-disk/service/repository/rpc/repository"
	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest) (resp *types.UserFileListResponse, err error) {
	//获得分页的初始下标和每页大小
	pageSize := req.Size
	if req.Size == 0 {
		pageSize = utils.DefaultPageSize
	}
	startPage := req.Page
	if startPage == 0 {
		startPage = 1
	}
	startIndex := pageSize * (startPage - 1)
	//根据文件夹identity，获取id，然后作为父目录id去搜目录下的数据
	userRepositoryInfo, err := l.svcCtx.UserRepositoryModel.FindByIdentity(l.ctx, req.Identity)
	if err != nil {
		return nil, errors.New("文件夹identity有误！")
	}
	allUserRepository, err := l.svcCtx.UserRepositoryModel.FindAllInPage(l.ctx, userRepositoryInfo.Id, userRepositoryInfo.UserIdentity.String, startIndex, pageSize)
	if err != nil {
		return nil, errors.New("该文件夹下搜索文件失败！")
	}
	newList := make([]*types.UserFile, 0)
	for _, userRepository := range allUserRepository {
		repositoryInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: userRepository.RepositoryIdentity.String})
		if err != nil {
			return nil, err
		}
		newList = append(newList, &types.UserFile{
			Id:                 userRepository.Id,
			Identity:           userRepository.Identity.String,
			RepositoryIdentity: userRepository.RepositoryIdentity.String,
			Name:               userRepository.Name.String,
			Ext:                repositoryInfo.Ext,
			Path:               repositoryInfo.Path,
			Size:               repositoryInfo.Size,
		})
	}
	return &types.UserFileListResponse{
		List:  newList,
		Count: int64(len(allUserRepository)),
	}, err
}
