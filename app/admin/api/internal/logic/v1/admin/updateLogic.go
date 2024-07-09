package admin

import (
	"blog/common/helper"
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

func (l *UpdateLogic) Update(req *types.AdminUpdateReq) error {
	password, _ := helper.PasswordHash(req.Password)
	data := model.Admin{
		ID:        req.Id,
		Name:      req.Name,
		Password:  password,
		Phone:     req.Phone,
		RoleID:    req.RoleId,
		CreatedAt: time.Now(),
	}
	err := l.svcCtx.DB.WithContext(l.ctx).Save(&data).Error
	if err != nil {
		errx.NewCodeError(errx.UpdateError)
	}
	return nil
}
