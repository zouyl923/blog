package permission

import (
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"time"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.SystemPermissionUpdateReq) error {
	data := model.AdminPermission{
		ID:        req.Id,
		MenuID:    req.MenuId,
		URI:       req.Uri,
		CreatedAt: time.Now(),
	}
	err := l.svcCtx.DB.WithContext(l.ctx).Debug().Where("menu_id =? ", req.MenuId).Save(&data).Error
	if err != nil {
		return errx.NewCodeError(errx.UpdateError)
	}
	return nil
}
