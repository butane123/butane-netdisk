package handler

import (
	"cloud-disk/common/response"
	"net/http"

	"cloud-disk/service/user_repository/api/internal/logic"
	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFolderCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderCreate(&req)
		response.Response(w, resp, err)
	}
}
