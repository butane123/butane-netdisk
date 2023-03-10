// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: repository.proto

package repository

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryClient interface {
	GetRepositoryPoolByRepositoryId(ctx context.Context, in *RepositoryReq, opts ...grpc.CallOption) (*RepositoryResp, error)
	DeleteById(ctx context.Context, in *DeleteByIdReq, opts ...grpc.CallOption) (*DeleteByIdResp, error)
}

type repositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryClient(cc grpc.ClientConnInterface) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) GetRepositoryPoolByRepositoryId(ctx context.Context, in *RepositoryReq, opts ...grpc.CallOption) (*RepositoryResp, error) {
	out := new(RepositoryResp)
	err := c.cc.Invoke(ctx, "/repository.repository/getRepositoryPoolByRepositoryId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) DeleteById(ctx context.Context, in *DeleteByIdReq, opts ...grpc.CallOption) (*DeleteByIdResp, error) {
	out := new(DeleteByIdResp)
	err := c.cc.Invoke(ctx, "/repository.repository/deleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
// All implementations must embed UnimplementedRepositoryServer
// for forward compatibility
type RepositoryServer interface {
	GetRepositoryPoolByRepositoryId(context.Context, *RepositoryReq) (*RepositoryResp, error)
	DeleteById(context.Context, *DeleteByIdReq) (*DeleteByIdResp, error)
	mustEmbedUnimplementedRepositoryServer()
}

// UnimplementedRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryServer struct {
}

func (UnimplementedRepositoryServer) GetRepositoryPoolByRepositoryId(context.Context, *RepositoryReq) (*RepositoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoryPoolByRepositoryId not implemented")
}
func (UnimplementedRepositoryServer) DeleteById(context.Context, *DeleteByIdReq) (*DeleteByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedRepositoryServer) mustEmbedUnimplementedRepositoryServer() {}

// UnsafeRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServer will
// result in compilation errors.
type UnsafeRepositoryServer interface {
	mustEmbedUnimplementedRepositoryServer()
}

func RegisterRepositoryServer(s grpc.ServiceRegistrar, srv RepositoryServer) {
	s.RegisterService(&Repository_ServiceDesc, srv)
}

func _Repository_GetRepositoryPoolByRepositoryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetRepositoryPoolByRepositoryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.repository/getRepositoryPoolByRepositoryId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetRepositoryPoolByRepositoryId(ctx, req.(*RepositoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.repository/deleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).DeleteById(ctx, req.(*DeleteByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Repository_ServiceDesc is the grpc.ServiceDesc for Repository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "repository.repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRepositoryPoolByRepositoryId",
			Handler:    _Repository_GetRepositoryPoolByRepositoryId_Handler,
		},
		{
			MethodName: "deleteById",
			Handler:    _Repository_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository.proto",
}
