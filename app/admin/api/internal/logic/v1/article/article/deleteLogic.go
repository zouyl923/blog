package article

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

func (l *DeleteLogic) Delete(req *types.ArticleDeleteReq) error {
	ids := make([]string, 0)
	if len(req.Uuid) > 0 {
		ids = append(ids, req.Uuid)
	}
	if len(req.Uuids) > 0 {
		ids = req.Uuids
	}
	var info []model.Article
	err := l.svcCtx.DB.WithContext(l.ctx).Debug().
		Where("uuid in (?)", ids).
		Find(&info).
		Update("is_del", 1).Error
	if err != nil {
		return errx.NewCodeError(errx.DeleteError)
	}
	return nil
}
