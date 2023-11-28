package jurisdiction

import (
	"context"
	"goDemo/internal/svc"
	"net/http"
)

type AuthorizationMiddleware struct {
	Ctx *svc.ServiceContext
}

func (m AuthorizationMiddleware) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("authorization")
		var userId string

		if len(authorization) != 0 {
			claims, err := m.Ctx.AuthorizationService.ParseTokenString(authorization)

			if err != nil {
				userId = claims.ID
			}
		}

		ctx := context.WithValue(r.Context(), "jti", userId)

		next(w, r.WithContext(ctx))
	}
}
