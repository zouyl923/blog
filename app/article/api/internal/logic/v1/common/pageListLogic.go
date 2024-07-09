package common

import (
	"blog/app/article/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListLogic {
	return &PageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageListLogic) PageList(req *types.SearchReq) (resp *types.PageList, err error) {
	list, err := l.svcCtx.ArticleCommonRpc.PageList(l.ctx, &rpc.SearchReq{
		Page:       req.Page,
		Keyword:    req.Keyword,
		CategoryId: req.CategoryId,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	var cList types.PageList
	helper.ExchangeStruct(list, &cList)
	return &cList, nil
}
