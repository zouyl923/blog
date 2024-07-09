package user

import (
	"blog/app/article/rpc/pb/rpc"
	"context"
	"strconv"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

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

func (l *DeleteLogic) Delete(req *types.DeleteReq) error {
	uId := l.ctx.Value("userId").(string)
	userId, err := strconv.ParseInt(uId, 10, 64)
	_, err = l.svcCtx.ArticleUserRpc.Delete(l.ctx, &rpc.DeleteReq{
		Uuid:   req.Uuid,
		UserId: userId,
	})
	return err
}
