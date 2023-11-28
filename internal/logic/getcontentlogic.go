package logic

import (
	"context"

	"goDemo/internal/svc"
	"goDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetContentLogic) GetContent(req *types.GetContentReq) error {
	l.svcCtx.ContentService.GetContentItem("")

	return nil
}
