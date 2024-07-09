package userservicelogic

import (
	"blog/app/article/rpc/internal/contants"
	"blog/database/model"
	"context"
	"time"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *rpc.DeleteReq) (*rpc.EmptyRes, error) {
	info := model.Article{}
	err := l.svcCtx.DB.WithContext(l.ctx).
		Where("uuid = ? ", in.Uuid).
		Where("user_id = ? ", in.UserId).
		First(info).
		Updates(&model.Article{
			IsDel:     1,
			DeletedAt: time.Now().Unix(),
		}).Error
	if err != nil {
		return nil, err
	}
	//清除缓存
	l.svcCtx.Cache.Del(contants.ArticleInfoKey + in.Uuid)
	return &rpc.EmptyRes{}, nil
}
