package user

import (
	"blog/app/user/rpc/client/userservice"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"

	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MainLogic {
	return &MainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MainLogic) Main(req *types.UserInfoReq) (resp *types.User, err error) {
	info, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.InfoReq{
		UserId: req.Id,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	cInfo := types.User{}
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
