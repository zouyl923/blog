package permission

import (
	"blog/common/response/errx"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.SystemPermissionDeleteReq) error {
	ids := make([]int64, 0)
	if req.Id > 0 {
		ids = append(ids, req.Id)
	}
	if len(req.Ids) > 0 {
		ids = req.Ids
	}
	var info []model.AdminPermission
	err := l.svcCtx.DB.WithContext(l.ctx).
		Where("id in (?)", ids).
		Delete(&info).Error
	if err != nil {
		return errx.NewCodeError(errx.DeleteError)
	}
	return nil
}
