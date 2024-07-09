package role

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SystemRoleSearchReq) (resp *types.SystemRolePageList, err error) {
	page := req.Page
	pageSize := req.PageSize
	offset := (page - 1) * pageSize
	var roles []model.AdminRole
	var total int64
	model := l.svcCtx.DB.WithContext(l.ctx)
	if len(req.Keyword) > 0 {
		model = model.Where("  name like ?  ", "%"+req.Keyword+"%")
	}
	model.
		Where("is_del = ?", 0).
		Offset(offset).
		Limit(pageSize).
		Find(&roles).
		Count(&total)
	//数据格式转换
	var data []types.SystemRole
	helper.Swap(roles, &data)
	resp = new(types.SystemRolePageList)
	resp.Page = page
	resp.PageSize = pageSize
	resp.Total = total
	resp.Data = data
	return resp, nil
}
