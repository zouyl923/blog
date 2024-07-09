package user

import (
	"blog/app/article/rpc/client/userservice"
	"context"
	"strconv"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	return &PushLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushLogic) Push(req *types.UpdateReq) error {
	uId := l.ctx.Value("userId").(string)
	userId, err := strconv.ParseInt(uId, 10, 64)
	_, err = l.svcCtx.ArticleUserRpc.Push(l.ctx, &userservice.UpdateReq{
		Uuid:       req.Uuid,
		Title:      req.Title,
		CategoryId: req.CategoryId,
		Cover:      req.Cover,
		Content:    req.Content,
		UserId:     userId,
	})
	return err
}
