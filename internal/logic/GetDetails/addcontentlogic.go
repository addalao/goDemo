package GetDetails

import (
	"context"
	"errors"
	"goDemo/internal/svc"
	"goDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContentLogic {
	return &AddContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddContentLogic) AddContent(req *types.AddContentReq) error {

	userId := l.svcCtx.UserService.GetId(l.ctx)

	if len(userId) == 0 {
		return errors.New("用户权限不足")
	}

	err := l.svcCtx.ContentService.Create(userId, req.Title, req.Content, "", req.Gender)

	if err != nil {
		return err
	}

	return nil

}
