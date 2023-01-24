package handler

import (
	"net/http"

	"cloud-disk/service/user_repository/api/internal/logic"
	"cloud-disk/service/user_repository/api/internal/svc"
	"cloud-disk/service/user_repository/api/internal/types"
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
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
