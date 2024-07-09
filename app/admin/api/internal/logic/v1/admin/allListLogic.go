package admin

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

func (l *AllListLogic) AllList(req *types.AdminSearchReq) (resp []types.AdminInfo, err error) {
	var list []model.Admin
	l.svcCtx.DB.WithContext(l.ctx).
		Where("is_del = ?", 0).
		Find(&list)
	//数据格式转换
	var cList []types.AdminInfo
	helper.Swap(list, &cList)
	return cList, nil
}
