package role

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

func (l *InfoLogic) Info(req *types.SystemRoleInfoReq) (resp *types.SystemRole, err error) {
	adminRole := model.AdminRole{}
	err = l.svcCtx.DB.WithContext(l.ctx).Where("id =?", req.Id).First(&adminRole).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.NotFundError)
	}

	var permissions []model.AdminRolePermission
	l.svcCtx.DB.WithContext(l.ctx).Where("role_id=?", adminRole.ID).Find(&permissions)
	var per []int32
	for _, v := range permissions {
		per = append(per, v.MenuID)
	}
	adminRoleInfo := new(types.SystemRole)
	helper.Swap(adminRole, &adminRoleInfo)
	adminRoleInfo.Permission = per
	return adminRoleInfo, nil
}
