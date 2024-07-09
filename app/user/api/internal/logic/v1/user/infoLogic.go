package user

import (
	"blog/app/user/rpc/client/userservice"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"
	"strconv"

	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"

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

func (l *InfoLogic) Info() (resp *types.User, err error) {
	userId := l.ctx.Value("userId").(string)
	uId, err := strconv.ParseInt(userId, 10, 64)
	info, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.InfoReq{
		UserId: uId,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	cInfo := types.User{}
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
