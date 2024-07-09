package user

import (
	"blog/app/article/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"
	"strconv"

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
	uId := l.ctx.Value("userId").(string)
	userId, err := strconv.ParseInt(uId, 10, 64)
	info, err := l.svcCtx.ArticleUserRpc.Info(l.ctx, &rpc.InfoReq{
		Uuid:   req.Uuid,
		UserId: userId,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	var cInfo types.Article
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
