package handler

import (
	"net/http"

	"cloud-disk/service/repository/api/internal/logic"
	"cloud-disk/service/repository/api/internal/svc"
	"cloud-disk/service/repository/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadByChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadByChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewFileUploadByChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadByChunk(&req, file, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}