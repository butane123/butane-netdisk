// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"butane-netdisk/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddVolumeReq       = user.AddVolumeReq
	AddVolumeResp      = user.AddVolumeResp
	DecreaseVolumeReq  = user.DecreaseVolumeReq
	DecreaseVolumeResp = user.DecreaseVolumeResp
	FindVolumeReq      = user.FindVolumeReq
	FindVolumeResp     = user.FindVolumeResp

	User interface {
		DecreaseVolume(ctx context.Context, in *DecreaseVolumeReq, opts ...grpc.CallOption) (*DecreaseVolumeResp, error)
		FindVolumeById(ctx context.Context, in *FindVolumeReq, opts ...grpc.CallOption) (*FindVolumeResp, error)
		AddVolume(ctx context.Context, in *AddVolumeReq, opts ...grpc.CallOption) (*AddVolumeResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) DecreaseVolume(ctx context.Context, in *DecreaseVolumeReq, opts ...grpc.CallOption) (*DecreaseVolumeResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DecreaseVolume(ctx, in, opts...)
}

func (m *defaultUser) FindVolumeById(ctx context.Context, in *FindVolumeReq, opts ...grpc.CallOption) (*FindVolumeResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FindVolumeById(ctx, in, opts...)
}

func (m *defaultUser) AddVolume(ctx context.Context, in *AddVolumeReq, opts ...grpc.CallOption) (*AddVolumeResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddVolume(ctx, in, opts...)
}
