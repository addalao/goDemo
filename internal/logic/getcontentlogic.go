package logic

import (
	"context"
	"strconv"

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

func (l *GetContentLogic) GetContent(req *types.GetContentReq) (resp *types.GetContentResp, err error) {
	getItems, total, err := l.svcCtx.ContentService.GetContentItem(req.Page, req.PageSize)

	items := make([]types.GetContentRespItem, 0)

	for _, item := range getItems {

		var (
			viewCount int64
			likeCount int64
		)

		id := strconv.FormatUint(uint64(item.ID), 10)

		viewCount, err = l.svcCtx.ViewService.GetCount(id)

		if err != nil {
			return nil, err
		}

		likeCount, err = l.svcCtx.LikeService.GetCount(id)

		if err != nil {
			return nil, err
		}

		items = append(items, types.GetContentRespItem{
			Id:      strconv.FormatUint(uint64(item.ID), 10),
			UserId:  item.UserId,
			Content: item.Content,
			Title:   item.Title,
			View:    viewCount,
			Like:    likeCount,
		})

	}

	if err != nil {
		return nil, err
	}

	return &types.GetContentResp{
		Items: items,
		Total: total,
	}, err
}
