package handler

import (
	"butane-netdisk/common/response"
	"net/http"

	"butane-netdisk/service/user/api/internal/logic"
	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		response.Response(w, resp, err)
	}
}
