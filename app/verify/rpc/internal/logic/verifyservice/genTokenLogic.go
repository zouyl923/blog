package verifyservicelogic

import (
	"blog/common/helper"
	"context"
	"time"

	"blog/app/verify/rpc/internal/svc"
	"blog/app/verify/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenTokenLogic) GenToken(in *rpc.GenTokenReq) (*rpc.GenTokenRes, error) {
	key := in.Server + ":" + in.Platform + ":" + in.Key
	tokenKey := "token:" + key
	refreshTokenKey := "refreshToken:" + key
	//生成token
	token, err := helper.GenToken(in.Key, l.svcCtx.Config.JwtSecret, 2*60*60)
	if err != nil {
		return nil, err
	}
	//缓存token
	l.svcCtx.Cache.Setex(tokenKey, token, int(in.Expire))
	//生成refreshToken
	str := in.Key + time.Now().String() + l.svcCtx.Config.JwtSecret
	refreshToken := helper.Hash256(str)
	//缓存token 用来刷新的
	l.svcCtx.Cache.Setex(refreshTokenKey, refreshToken, 30*60*60)
	return &rpc.GenTokenRes{
		Server:       in.Server,
		Platform:     in.Platform,
		Key:          in.Key,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
