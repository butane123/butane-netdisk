// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-disk/service/share/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/share/basic/detail",
				Handler: ShareBasicDetailHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/share/basic/create",
				Handler: ShareBasicCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share/basic/save",
				Handler: ShareBasicSaveHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
