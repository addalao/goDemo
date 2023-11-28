package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"goDemo/internal/svc"
)

type GetContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentLogic) GetContent() error {
	// todo: add your logic here and delete this line

	return nil
}
