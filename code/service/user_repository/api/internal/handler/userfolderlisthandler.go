package handler

import (
	"butane-netdisk/common/response"
	"net/http"

	"butane-netdisk/service/user_repository/api/internal/logic"
	"butane-netdisk/service/user_repository/api/internal/svc"
	"butane-netdisk/service/user_repository/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFolderListLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderList(&req)
		response.Response(w, resp, err)
	}
}
