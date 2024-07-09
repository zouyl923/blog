package verifyservicelogic

import (
	"context"

	"blog/app/verify/rpc/internal/svc"
	"blog/app/verify/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTokenLogic {
	return &RemoveTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveTokenLogic) RemoveToken(in *rpc.RemoveTokenReq) (*rpc.EmptyRes, error) {
	key := in.Server + ":" + in.Platform + ":" + in.Key
	tokenKey := "token:" + key
	refreshTokenKey := "refreshToken:" + key
	//删除缓存
	l.svcCtx.Cache.Del(tokenKey)
	l.svcCtx.Cache.Del(refreshTokenKey)
	return &rpc.EmptyRes{}, nil
}
