package logic

import (
	"butane-netdisk/common/utils"
	"butane-netdisk/service/share/model"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"
	"context"
	"encoding/json"
	"fmt"

	"butane-netdisk/service/share/api/internal/svc"
	"butane-netdisk/service/share/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest) (resp *types.ShareBasicCreateResponse, err error) {
	RepositoryIdInfo, err := l.svcCtx.UserRepositoryRpc.FindRepositoryIdById(l.ctx, &userRepository.FindRepositoryIdReq{Id: req.UserRepositoryId})
	if err != nil {
		return nil, err
	}
	newId := utils.GenerateNewId(l.svcCtx.RedisClient, "share")
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ShareBasicModel.InsertWithId(l.ctx, &model.ShareBasic{
		Id:               newId,
		UserId:           userId,
		RepositoryId:     RepositoryIdInfo.RepositoryId,
		UserRepositoryId: req.UserRepositoryId,
		ExpiredTime:      req.ExpiredTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicCreateResponse{Id: newId}, nil
}
