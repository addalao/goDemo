// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LogingRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type GetDetailsRes struct {
	UserName string `json:"userName"`
}

type GetContentReq struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"pageSize"`
}
