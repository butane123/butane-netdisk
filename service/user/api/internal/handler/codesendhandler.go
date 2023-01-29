package handler

import (
	"cloud-disk/common/response"
	"net/http"

	"cloud-disk/service/user/api/internal/logic"
	"cloud-disk/service/user/api/internal/svc"
	"cloud-disk/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.CodeSend(&req)
		response.Response(w, resp, err)
	}
}
