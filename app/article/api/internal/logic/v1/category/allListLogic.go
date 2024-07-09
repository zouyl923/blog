package category

import (
	"blog/app/article/rpc/pb/rpc"
	"blog/common/helper"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllListLogic {
	return &AllListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllListLogic) AllList() (resp []types.Category, err error) {
	list, err := l.svcCtx.ArticleCommonRpc.CategoryList(l.ctx, &rpc.EmptyReq{})
	var cList []types.Category
	helper.ExchangeStruct(list.List, &cList)
	return cList, nil
}
