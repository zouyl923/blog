package role

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

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

func (l *AllListLogic) AllList() (resp []types.SystemRole, err error) {
	var list []model.AdminRole
	l.svcCtx.DB.WithContext(l.ctx).
		Where("is_del = ?", 0).
		Where("is_hid = ?", 0).
		Find(&list)
	var cList []types.SystemRole
	helper.Swap(list, &cList)
	return cList, nil
}
