package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goDemo/internal/logic"
	"goDemo/internal/svc"
)

func GetContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetContentLogic(r.Context(), svcCtx)
		err := l.GetContent()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
