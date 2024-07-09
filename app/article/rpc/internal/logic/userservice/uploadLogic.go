package userservicelogic

import (
	"context"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *rpc.UploadReq) (*rpc.UploadRes, error) {
	// todo: add your logic here and delete this line

	return &rpc.UploadRes{}, nil
}
