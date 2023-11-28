package GetDetails

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goDemo/internal/logic/GetDetails"
	"goDemo/internal/svc"
	"goDemo/internal/types"
)

func AddContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddContentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := GetDetails.NewAddContentLogic(r.Context(), svcCtx)
		err := l.AddContent(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
