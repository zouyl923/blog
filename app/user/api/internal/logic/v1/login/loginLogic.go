package login

import (
	"blog/app/user/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"

	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	loginRes, err := l.svcCtx.UserRpc.Login(l.ctx, &rpc.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	cUser := types.User{}
	helper.ExchangeStruct(loginRes.User, &cUser)
	resp = new(types.LoginRes)
	resp.User = cUser
	resp.Token = loginRes.Token
	resp.RefreshToken = loginRes.RefreshToken
	return resp, nil
}
