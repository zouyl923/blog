package article

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.ArticleSearchReq) (resp *types.ArticlePageList, err error) {
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize
	var list []model.Article
	var total int64
	model := l.svcCtx.DB.WithContext(l.ctx)
	if len(req.Keyword) > 0 {
		model = model.Where(" ( title like  ? ) ", "%"+req.Keyword+"%")
	}
	model.
		Where("is_del = ?", 0).
		Preload("CategoryInfo").
		Preload("DetailInfo").
		Offset(offset).
		Limit(pageSize).
		Find(&list).
		Count(&total)
	//数据格式转换
	var cList []types.Article
	helper.Swap(list, &cList)
	resp = new(types.ArticlePageList)
	resp.Page = page
	resp.PageSize = pageSize
	resp.Total = total
	resp.Data = cList
	return resp, nil
}
