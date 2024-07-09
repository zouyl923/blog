package common

import (
	"blog/app/article/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

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

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.Article, err error) {
	info, err := l.svcCtx.ArticleCommonRpc.Info(l.ctx, &rpc.InfoReq{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	var cInfo types.Article
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
