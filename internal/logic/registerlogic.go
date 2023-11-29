package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"goDemo/internal/svc"
	"goDemo/internal/types"
	"goDemo/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {

	tx := l.svcCtx.Db.First(&models.User{}, "username = ?", req.UserName)

	if tx.RowsAffected != 0 {
		return fmt.Errorf("已有相同用户")
	}

	// 生成哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	l.svcCtx.Db.Create(&models.User{
		Username: req.UserName,
		Password: string(hashedPassword),
		Gender:   req.Gender,
	})
	return nil
}
