package commonservicelogic

import (
	"blog/app/article/rpc/internal/contants"
	"blog/common/helper"
	"blog/database/model"
	"context"
	"encoding/json"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryListLogic) CategoryList(in *rpc.EmptyReq) (*rpc.CategoryRes, error) {
	var list []model.ArticleCategory
	var cList []*rpc.Category
	cache, _ := l.svcCtx.Cache.Get(contants.CategoryCacheKey)
	_ = json.Unmarshal([]byte(cache), &list)
	if len(list) < 1 {
		l.svcCtx.DB.WithContext(l.ctx).Debug().
			Where("is_del = ? ", 0).
			Where("is_hid = ? ", 0).
			Find(&list)
		jList, _ := json.Marshal(list)
		l.svcCtx.Cache.Setex(contants.CategoryCacheKey, string(jList), 20*60*00)
	}
	helper.ExchangeStruct(list, &cList)
	return &rpc.CategoryRes{List: cList}, nil
}
