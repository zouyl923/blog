package verifyservicelogic

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"context"

	"blog/app/verify/rpc/internal/svc"
	"blog/app/verify/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *rpc.AuthReq) (*rpc.AuthRes, error) {
	//解析token
	claims, err := helper.ParseToken(in.Token, l.svcCtx.Config.JwtSecret)
	if err != nil {
		return nil, err
	}
	key := in.Server + in.Platform + ":" + ":" + claims.Key
	//验证token是否有效
	tokenKey := "token:" + key
	token, err := l.svcCtx.Cache.Get(tokenKey)
	if err != nil {
		return nil, err
	}
	//验证token是否有效
	refreshTokenKey := "refreshToken:" + key
	refreshToken, err := l.svcCtx.Cache.Get(refreshTokenKey)
	if err != nil {
		return nil, err
	}

	if len(token) == 0 || len(refreshToken) == 0 {
		return nil, errx.NewCodeError(errx.LoginError)
	}

	var jwtMap rpc.JwtMap
	helper.ExchangeStruct(claims, &jwtMap)
	return &rpc.AuthRes{
		Server:       in.Server,
		Key:          claims.Key,
		Token:        token,
		RefreshToken: refreshToken,
		JwtMap:       &jwtMap,
	}, nil
}
