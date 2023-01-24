// Code generated by goctl. DO NOT EDIT.
// Source: userRepository.proto

package userrepositoryclient

import (
	"context"

	"cloud-disk/service/user_repository/rpc/userRepository"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateByShareReply      = userRepository.CreateByShareReply
	CreateByShareReq        = userRepository.CreateByShareReq
	FindRepositoryIdReply   = userRepository.FindRepositoryIdReply
	FindRepositoryIdReq     = userRepository.FindRepositoryIdReq
	RepositoryIdReq         = userRepository.RepositoryIdReq
	UserRepositoryNameReply = userRepository.UserRepositoryNameReply

	UserRepository interface {
		GetUserRepositoryNameByRepositoryId(ctx context.Context, in *RepositoryIdReq, opts ...grpc.CallOption) (*UserRepositoryNameReply, error)
		FindRepositoryIdentityByIdentity(ctx context.Context, in *FindRepositoryIdReq, opts ...grpc.CallOption) (*FindRepositoryIdReply, error)
		CreateByShare(ctx context.Context, in *CreateByShareReq, opts ...grpc.CallOption) (*CreateByShareReply, error)
	}

	defaultUserRepository struct {
		cli zrpc.Client
	}
)

func NewUserRepository(cli zrpc.Client) UserRepository {
	return &defaultUserRepository{
		cli: cli,
	}
}

func (m *defaultUserRepository) GetUserRepositoryNameByRepositoryId(ctx context.Context, in *RepositoryIdReq, opts ...grpc.CallOption) (*UserRepositoryNameReply, error) {
	client := userRepository.NewUserRepositoryClient(m.cli.Conn())
	return client.GetUserRepositoryNameByRepositoryId(ctx, in, opts...)
}

func (m *defaultUserRepository) FindRepositoryIdentityByIdentity(ctx context.Context, in *FindRepositoryIdReq, opts ...grpc.CallOption) (*FindRepositoryIdReply, error) {
	client := userRepository.NewUserRepositoryClient(m.cli.Conn())
	return client.FindRepositoryIdentityByIdentity(ctx, in, opts...)
}

func (m *defaultUserRepository) CreateByShare(ctx context.Context, in *CreateByShareReq, opts ...grpc.CallOption) (*CreateByShareReply, error) {
	client := userRepository.NewUserRepositoryClient(m.cli.Conn())
	return client.CreateByShare(ctx, in, opts...)
}