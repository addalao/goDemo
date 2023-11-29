package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goDemo/internal/logic"
	"goDemo/internal/svc"
	"goDemo/internal/types"
)

func GetContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetContentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetContentLogic(r.Context(), svcCtx)
		resp, err := l.GetContent(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
