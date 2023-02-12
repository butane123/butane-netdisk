package logic

import (
	"butane-netdisk/common/errorx"
	"butane-netdisk/common/utils"
	"butane-netdisk/service/repository/rpc/types/repository"
	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"
	"context"
	"encoding/json"
	"fmt"

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
	//根据文件夹id，然后作为父目录id去搜目录下的数据
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	allUserRepository, err := l.svcCtx.UserRepositoryModel.FindAllInPage(l.ctx, req.Id, userId, startIndex, pageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("该文件夹下搜索文件失败！")
	}
	newList := make([]*types.UserFile, 0)
	for _, userRepository := range allUserRepository {
		repositoryInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: userRepository.RepositoryId})
		if err != nil {
			return nil, err
		}
		newList = append(newList, &types.UserFile{
			Id:           userRepository.Id,
			RepositoryId: userRepository.RepositoryId,
			Name:         userRepository.Name,
			Ext:          repositoryInfo.Ext,
			Path:         repositoryInfo.Path,
			Size:         repositoryInfo.Size,
		})
	}
	return &types.UserFileListResponse{
		List:  newList,
		Count: int64(len(allUserRepository)),
	}, err
}
