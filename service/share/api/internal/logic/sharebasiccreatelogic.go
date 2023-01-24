package logic

import (
	"cloud-disk/common/utils"
	"cloud-disk/service/share/model"
	"cloud-disk/service/user_repository/rpc/userRepository"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"cloud-disk/service/share/api/internal/svc"
	"cloud-disk/service/share/api/internal/types"

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
	RepositoryIdentityInfo, err := l.svcCtx.UserRepositoryRpc.FindRepositoryIdentityByIdentity(l.ctx, &userRepository.FindRepositoryIdReq{Identity: req.UserRepositoryIdentity})
	if err != nil {
		return nil, err
	}
	newIdentity := utils.GenerateUUID()
	_, err = l.svcCtx.ShareBasicModel.Insert(l.ctx, &model.ShareBasic{
		Identity:               sql.NullString{newIdentity, true},
		UserIdentity:           sql.NullString{json.Number(fmt.Sprintf("%v", l.ctx.Value("userIdentity"))).String(), true},
		RepositoryIdentity:     sql.NullString{RepositoryIdentityInfo.RepositoryIdentity, true},
		UserRepositoryIdentity: sql.NullString{req.UserRepositoryIdentity, true},
		ExpiredTime:            sql.NullInt64{req.ExpiredTime, true},
	})
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicCreateResponse{Identity: newIdentity}, nil
}
