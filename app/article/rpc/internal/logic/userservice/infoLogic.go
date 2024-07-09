package userservicelogic

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *rpc.InfoReq) (*rpc.Article, error) {
	info := model.Article{}
	md := l.svcCtx.DB.WithContext(l.ctx)
	err := md.
		Where("uuid = ?", in.Uuid).
		Where("is_del = ?", 0).
		Where("user_id = ?", in.UserId).
		Preload("DetailInfo").
		First(&info).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.NotFundError)
	}
	cInfo := rpc.Article{}
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
