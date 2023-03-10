// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"butane-netdisk/service/user/rpc/internal/logic"
	"butane-netdisk/service/user/rpc/internal/svc"
	"butane-netdisk/service/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) DecreaseVolume(ctx context.Context, in *user.DecreaseVolumeReq) (*user.DecreaseVolumeResp, error) {
	l := logic.NewDecreaseVolumeLogic(ctx, s.svcCtx)
	return l.DecreaseVolume(in)
}

func (s *UserServer) FindVolumeById(ctx context.Context, in *user.FindVolumeReq) (*user.FindVolumeResp, error) {
	l := logic.NewFindVolumeByIdLogic(ctx, s.svcCtx)
	return l.FindVolumeById(in)
}

func (s *UserServer) AddVolume(ctx context.Context, in *user.AddVolumeReq) (*user.AddVolumeResp, error) {
	l := logic.NewAddVolumeLogic(ctx, s.svcCtx)
	return l.AddVolume(in)
}
