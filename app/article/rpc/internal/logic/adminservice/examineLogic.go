package adminservicelogic

import (
	"context"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExamineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamineLogic {
	return &ExamineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExamineLogic) Examine(in *rpc.DeleteReq) (*rpc.EmptyRes, error) {
	// todo: add your logic here and delete this line

	return &rpc.EmptyRes{}, nil
}
