// Code generated by goctl. DO NOT EDIT.
// Source: userRepository.proto

package server

import (
	"context"

	"butane-netdisk/service/user_repository/rpc/internal/logic"
	"butane-netdisk/service/user_repository/rpc/internal/svc"
	"butane-netdisk/service/user_repository/rpc/types/userRepository"
)

type UserRepositoryServer struct {
	svcCtx *svc.ServiceContext
	userRepository.UnimplementedUserRepositoryServer
}

func NewUserRepositoryServer(svcCtx *svc.ServiceContext) *UserRepositoryServer {
	return &UserRepositoryServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRepositoryServer) GetUserRepositoryNameByRepositoryId(ctx context.Context, in *userRepository.RepositoryIdReq) (*userRepository.UserRepositoryNameReply, error) {
	l := logic.NewGetUserRepositoryNameByRepositoryIdLogic(ctx, s.svcCtx)
	return l.GetUserRepositoryNameByRepositoryId(in)
}

func (s *UserRepositoryServer) FindRepositoryIdById(ctx context.Context, in *userRepository.FindRepositoryIdReq) (*userRepository.FindRepositoryIdReply, error) {
	l := logic.NewFindRepositoryIdByIdLogic(ctx, s.svcCtx)
	return l.FindRepositoryIdById(in)
}

func (s *UserRepositoryServer) CreateByShare(ctx context.Context, in *userRepository.CreateByShareReq) (*userRepository.CreateByShareReply, error) {
	l := logic.NewCreateByShareLogic(ctx, s.svcCtx)
	return l.CreateByShare(in)
}