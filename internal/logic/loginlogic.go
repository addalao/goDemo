package logic

import (
	"context"
	"fmt"
	"goDemo/internal/svc"
	"goDemo/internal/types"
	"goDemo/models"
	"goDemo/tool"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LogingRes, err error) {
	user := models.User{}
	l.svcCtx.Db.First(&user, "UserName = ?", req.UserName)
	if user.ID == 0 {
		return nil, fmt.Errorf("用户名错误")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("密码错误")
	}
	token, err := tool.GenerateToken(user.ID, l.svcCtx.Config)
	if err != nil {
		return nil, err
	}
	return &types.LogingRes{
		Token: token,
	}, nil
}
