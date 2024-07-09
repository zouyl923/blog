package category

import (
	"blog/common/helper"
	"blog/common/static/cKey"
	"blog/database/model"
	"context"
	"encoding/json"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeLogic {
	return &TreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeLogic) Tree() (resp []types.ArticleCategory, err error) {
	var list []model.ArticleCategory
	var cList []types.ArticleCategory
	cache, _ := l.svcCtx.Cache.Get(cKey.ArticleCategoryTreeKey)
	_ = json.Unmarshal([]byte(cache), &cList)
	if len(list) < 1 {
		l.svcCtx.DB.WithContext(l.ctx).
			Where("is_del = ? ", 0).
			Where("is_hid = ? ", 0).
			Find(&list)
		helper.Swap(list, &cList)
		cList = GetTree(cList, 0)
		jList, _ := json.Marshal(cList)
		l.svcCtx.Cache.Setex(cKey.ArticleCategoryTreeKey, string(jList), cKey.ArticleCategoryTreeTtl)
	}
	return cList, nil
}

func GetTree(list []types.ArticleCategory, pid int64) []types.ArticleCategory {
	tree := make([]types.ArticleCategory, 0)
	for _, v := range list {
		if int64(v.ParentId) == pid {
			v.Children = GetTree(list, v.Id)
			tree = append(tree, v)
		}
	}
	return tree
}
