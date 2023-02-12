package main

import (
	"butane-netdisk/common/errorx"
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"butane-netdisk/service/repository/rpc/internal/config"
	"butane-netdisk/service/repository/rpc/internal/server"
	"butane-netdisk/service/repository/rpc/internal/svc"
	"butane-netdisk/service/repository/rpc/types/repository"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/repository.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		repository.RegisterRepositoryServer(grpcServer, server.NewRepositoryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
