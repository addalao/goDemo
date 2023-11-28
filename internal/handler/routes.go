// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	GetDetails "goDemo/internal/handler/GetDetails"
	"goDemo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/content",
				Handler: GetContentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/GetDetails",
				Handler: GetDetails.GetDetailsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/content/add",
				Handler: GetDetails.AddContentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
