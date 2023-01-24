package logic

import (
	"context"

	"cloud-disk/service/user_repository/rpc/internal/svc"
	"cloud-disk/service/user_repository/rpc/userRepository"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRepositoryIdentityByIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRepositoryIdentityByIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRepositoryIdentityByIdentityLogic {
	return &FindRepositoryIdentityByIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRepositoryIdentityByIdentityLogic) FindRepositoryIdentityByIdentity(in *userRepository.FindRepositoryIdReq) (*userRepository.FindRepositoryIdReply, error) {
	userRepositoryInfo, err := l.svcCtx.UserRepositoryModel.FindByIdentity(l.ctx, in.Identity)
	if err != nil {
		return nil, err
	}
	return &userRepository.FindRepositoryIdReply{RepositoryIdentity: userRepositoryInfo.RepositoryIdentity.String}, nil
}
