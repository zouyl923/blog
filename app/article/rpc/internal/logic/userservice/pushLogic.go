package userservicelogic

import (
	"blog/app/article/rpc/internal/contants"
	"blog/app/uuid/client/uuidservice"
	"blog/database/model"
	"context"
	"time"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	return &PushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushLogic) Push(in *rpc.UpdateReq) (*rpc.EmptyRes, error) {
	var uid string
	if len(in.Uuid) > 20 {
		uid = in.Uuid
	} else {
		biz, err := l.svcCtx.UuidRpc.GenUuid(l.ctx, &uuidservice.GenUuidReq{
			BizType: "article",
		})
		if err != nil {
			return nil, err
		}
		uid = biz.BizId
	}
	info := model.Article{
		Uuid:       uid,
		Title:      in.Title,
		Cover:      in.Cover,
		CategoryId: in.CategoryId,
		UserId:     int32(in.UserId),
		CreatedAt:  time.Now(),
	}
	detail := model.ArticleDetail{
		ArticleUuid: uid,
		Content:     in.Content,
	}
	md := l.svcCtx.DB.WithContext(l.ctx)
	md.Begin()
	err := md.Save(&info).Error
	if err != nil {
		md.Rollback()
		return nil, err
	}
	err = md.Save(&detail).Error
	if err != nil {
		md.Rollback()
		return nil, err
	}
	md.Commit()
	//清除缓存
	l.svcCtx.Cache.Del(contants.ArticleInfoKey + in.Uuid)
	return &rpc.EmptyRes{}, nil
}
