package main

import (
	"flag"
	"fmt"

	"cloud-disk/service/user_repository/rpc/internal/config"
	"cloud-disk/service/user_repository/rpc/internal/server"
	"cloud-disk/service/user_repository/rpc/internal/svc"
	"cloud-disk/service/user_repository/rpc/userRepository"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userrepository.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userRepository.RegisterUserRepositoryServer(grpcServer, server.NewUserRepositoryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
