package logic

import (
	"cloud-disk/service/repository/rpc/repository"
	"cloud-disk/service/user_repository/rpc/userRepository"
	"context"
	"errors"

	"cloud-disk/service/share/api/internal/svc"
	"cloud-disk/service/share/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	//1 增加点击数
	shareBasicInfo, err := l.svcCtx.ShareBasicModel.AddOneClick(l.ctx, req.Identity)
	if err != nil {
		return nil, errors.New("增加点击数失败！")
	}
	//2 连表查询，通过连repostitoryid和返回user库name、repositoryPool库3个值
	userRepositoryName, err := l.svcCtx.UserRepositoryRpc.GetUserRepositoryNameByRepositoryId(l.ctx, &userRepository.RepositoryIdReq{
		RepositoryId: shareBasicInfo.RepositoryIdentity.String,
	})
	if err != nil {
		return nil, errors.New("无法获得用户储存库的信息！")
	}
	RepositoryPool, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{
		RepositoryId: shareBasicInfo.RepositoryIdentity.String,
	})
	if err != nil {
		return nil, errors.New("无法获得储存池的信息！")
	}
	return &types.DetailResponse{
		RepositoryIdentity: shareBasicInfo.RepositoryIdentity.String,
		Name:               userRepositoryName.RepositoryName,
		Ext:                RepositoryPool.Ext,
		Size:               RepositoryPool.Size,
		Path:               RepositoryPool.Path,
	}, nil
}
