package logic

import (
	"cloud-disk/service/user_repository/rpc/internal/svc"
	"cloud-disk/service/user_repository/rpc/userRepository"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRepositoryNameByRepositoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRepositoryNameByRepositoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRepositoryNameByRepositoryIdLogic {
	return &GetUserRepositoryNameByRepositoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRepositoryNameByRepositoryIdLogic) GetUserRepositoryNameByRepositoryId(in *userRepository.RepositoryIdReq) (*userRepository.UserRepositoryNameReply, error) {
	userInfo, err := l.svcCtx.UserRepositoryModel.FindByRepositoryId(l.ctx, in.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &userRepository.UserRepositoryNameReply{RepositoryName: userInfo.Name.String}, nil
}
