package article

import (
	"blog/common/response/errx"
	uuid2 "blog/common/uuid"
	"blog/database/model"
	"context"
	"strconv"
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

func (l *UpdateLogic) Update(req *types.ArticleUpdateReq) error {
	uuid := req.Uuid
	if len(uuid) < 1 {
		snow, _ := uuid2.NewSnowFlake(1)
		uid := snow.Next()
		uuid = strconv.Itoa(int(uid))
	}
	data := &model.Article{
		Uuid:       uuid,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Cover:      req.Cover,
		CreatedAt:  time.Now(),
	}
	tx := l.svcCtx.DB.WithContext(l.ctx).Begin()
	err := tx.Save(data).Error
	if err != nil {
		tx.Rollback()
		return errx.NewCodeError(errx.UpdateError)
	}
	err = tx.Save(&types.ArticleDetail{
		ArticleUuid: uuid,
		Content:     req.Content,
	}).Error
	if err != nil {
		tx.Rollback()
		return errx.NewCodeError(errx.UpdateError)
	}
	tx.Commit()
	return nil
}
