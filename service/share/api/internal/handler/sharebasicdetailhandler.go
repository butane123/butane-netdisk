package handler

import (
	"cloud-disk/common/response"
	"net/http"

	"cloud-disk/service/share/api/internal/logic"
	"cloud-disk/service/share/api/internal/svc"
	"cloud-disk/service/share/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShareBasicDetailLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicDetail(&req)
		response.Response(w, resp, err)
	}
}
