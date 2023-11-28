package GetDetails

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goDemo/internal/logic/GetDetails"
	"goDemo/internal/svc"
)

func GetDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := GetDetails.NewGetDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetDetails()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
