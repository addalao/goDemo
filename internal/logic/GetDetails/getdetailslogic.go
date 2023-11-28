package GetDetails

import (
	"context"
	"fmt"
	"goDemo/internal/svc"
	"goDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailsLogic {
	return &GetDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailsLogic) GetDetails() (resp *types.GetDetailsRes, err error) {
	//tool.ParseToken(resp.)
	value := l.ctx.Value("authorization")
	fmt.Printf("这是Value--%v \n", value)
	return &types.GetDetailsRes{}, nil
}
