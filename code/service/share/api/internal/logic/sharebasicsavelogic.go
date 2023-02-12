package logic

import (
	"butane-netdisk/service/repository/rpc/types/repository"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/share/api/internal/svc"
	"butane-netdisk/service/share/api/internal/types"

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
	nameInfo, err := l.svcCtx.RepositoryRpc.GetRepositoryPoolByRepositoryId(l.ctx, &repository.RepositoryReq{RepositoryId: req.RepositoryId})
	if err != nil {
		return nil, err
	}
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	idInfo, err := l.svcCtx.UserRepositoryRpc.CreateByShare(l.ctx, &userRepository.CreateByShareReq{
		UserId:       userId,
		ParentId:     req.ParentId,
		RepositoryId: req.RepositoryId,
		Name:         nameInfo.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicSaveResponse{Id: idInfo.Id}, nil
}
