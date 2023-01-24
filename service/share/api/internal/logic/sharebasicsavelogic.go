package logic

import (
	"cloud-disk/service/user_repository/rpc/userRepository"
	"context"
	"encoding/json"
	"fmt"

	"cloud-disk/service/share/api/internal/svc"
	"cloud-disk/service/share/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest) (resp *types.ShareBasicSaveResponse, err error) {
	nameInfo, err := l.svcCtx.UserRepositoryRpc.GetUserRepositoryNameByRepositoryId(l.ctx, &userRepository.RepositoryIdReq{RepositoryId: req.RepositoryIdentity})
	if err != nil {
		return nil, err
	}
	identityInfo, err := l.svcCtx.UserRepositoryRpc.CreateByShare(l.ctx, &userRepository.CreateByShareReq{
		UserIdentity:       json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String(),
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               nameInfo.RepositoryName,
	})
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicSaveResponse{Identity: identityInfo.Identity}, nil
}
