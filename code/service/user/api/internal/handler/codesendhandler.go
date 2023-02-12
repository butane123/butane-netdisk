package handler

import (
	"butane-netdisk/common/response"
	"net/http"

	"butane-netdisk/service/user/api/internal/logic"
	"butane-netdisk/service/user/api/internal/svc"
	"butane-netdisk/service/user/api/internal/types"

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
