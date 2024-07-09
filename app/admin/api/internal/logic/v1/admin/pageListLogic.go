package admin

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

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

func (l *PageListLogic) PageList(req *types.AdminSearchReq) (resp *types.AdminPageList, err error) {
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize
	var list []model.Admin
	var total int64
	model := l.svcCtx.DB.WithContext(l.ctx)
	if len(req.Keyword) > 0 {
		model = model.Where(" ( name like ? or phone like ? ) ", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	model.
		Where("is_del = ?", 0).
		Preload("RoleInfo").
		Offset(offset).
		Limit(pageSize).
		Find(&list).
		Count(&total)
	//数据格式转换
	var cList []types.AdminInfo
	helper.Swap(list, &cList)
	resp = new(types.AdminPageList)
	resp.Page = page
	resp.PageSize = pageSize
	resp.Total = total
	resp.Data = cList
	return resp, nil
}
