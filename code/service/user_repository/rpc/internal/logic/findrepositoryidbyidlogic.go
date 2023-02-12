package logic

import (
	"context"

	"butane-netdisk/service/user_repository/rpc/internal/svc"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRepositoryIdByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRepositoryIdByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRepositoryIdByIdLogic {
	return &FindRepositoryIdByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRepositoryIdByIdLogic) FindRepositoryIdById(in *userRepository.FindRepositoryIdReq) (*userRepository.FindRepositoryIdReply, error) {
	userRepositoryInfo, err := l.svcCtx.UserRepositoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &userRepository.FindRepositoryIdReply{RepositoryId: userRepositoryInfo.RepositoryId}, nil
}
