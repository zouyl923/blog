package userservicelogic

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListLogic {
	return &PageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 个人文章
func (l *PageListLogic) PageList(in *rpc.SearchReq) (*rpc.PageData, error) {
	var list []model.Article
	var total int64
	var offset int
	page := int(in.Page)
	pageSize := 10
	offset = (page - 1) * pageSize
	md := l.svcCtx.DB.WithContext(l.ctx)
	if in.CategoryId > 0 {
		md = md.Where("category_id = ? ", in.CategoryId)
	}
	if len(in.Keyword) > 0 {
		md = md.Where("title like ?", "%"+in.Keyword+"%")
	}
	md.
		Where("user_id = ?", in.UserId).
		Where("is_del = ?", 0).
		Preload("CategoryInfo").
		Offset(offset).Limit(pageSize).
		Find(&list).Count(&total)
	var cList []*rpc.Article
	helper.ExchangeStruct(list, &cList)
	return &rpc.PageData{
		Total:    total,
		Page:     in.Page,
		PageSize: int32(pageSize),
		Data:     cList,
	}, nil
}
