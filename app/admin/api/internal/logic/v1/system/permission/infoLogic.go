package permission

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.SystemPermissionInfoReq) (resp *types.SystemPermission, err error) {
	info := model.AdminPermission{}
	err = l.svcCtx.DB.WithContext(l.ctx).
		Where("id =?", req.Id).
		First(&info).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.NotFundError)
	}
	menu := model.AdminMenu{}
	l.svcCtx.DB.WithContext(l.ctx).Where("id=?", info.MenuID).First(&menu)
	cInfo := new(types.SystemPermission)
	helper.Swap(info, cInfo)
	cInfo.Menu = append(cInfo.Menu, menu.ParentID, int32(menu.ID))
	return cInfo, nil
}
