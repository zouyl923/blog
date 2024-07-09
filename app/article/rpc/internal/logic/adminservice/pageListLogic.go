package adminservicelogic

import (
	"context"

	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListLogic {
	return &PageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理后台
func (l *PageListLogic) PageList(in *rpc.SearchReq) (*rpc.PageData, error) {
	// todo: add your logic here and delete this line

	return &rpc.PageData{}, nil
}
