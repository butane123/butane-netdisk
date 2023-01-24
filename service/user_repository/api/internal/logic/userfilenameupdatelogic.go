package logic

import (
	"context"
	"database/sql"

	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest) (resp *types.UserFileNameUpdateResponse, err error) {
	userFileInfo, err := l.svcCtx.UserRepositoryModel.FindByIdentity(l.ctx, req.Identity)
	if err != nil {
		return nil, err
	}
	userFileInfo.Name = sql.NullString{
		String: req.Name,
		Valid:  true,
	}
	err = l.svcCtx.UserRepositoryModel.Update(l.ctx, userFileInfo)
	if err != nil {
		return nil, err
	}
	return &types.UserFileNameUpdateResponse{}, nil
}
