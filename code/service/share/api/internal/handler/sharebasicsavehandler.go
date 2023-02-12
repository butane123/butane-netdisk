package handler

import (
	"butane-netdisk/common/response"
	"net/http"

	"butane-netdisk/service/share/api/internal/logic"
	"butane-netdisk/service/share/api/internal/svc"
	"butane-netdisk/service/share/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicSaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicSaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShareBasicSaveLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicSave(&req)
		response.Response(w, resp, err)
	}
}
