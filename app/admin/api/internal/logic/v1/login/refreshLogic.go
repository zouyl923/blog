package login

import (
	"blog/app/admin/api/internal/types"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/common/static/cKey"
	"context"
	"net/http"
	"time"

	"blog/app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新token
func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(r *http.Request) (resp *types.LoginRes, err error) {
	token := r.Header.Get("token")
	refreshToken := r.Header.Get("refresh-token")
	if token == "" || refreshToken == "" {
		return nil, errx.NewCodeError(errx.LoginError)
	}
	//解析token
	claims, err := helper.ParseToken(token, l.svcCtx.Config.JwtSecret)
	if err != nil {
		return nil, errx.NewCodeError(errx.LoginError)
	}
	adminId := claims.Key
	refreshTokenKey := cKey.AdminRefreshTokenKey + adminId
	cRefreshToken, _ := l.svcCtx.Cache.Get(refreshTokenKey)
	if cRefreshToken == "" || cRefreshToken != refreshToken {
		return nil, errx.NewCodeError(errx.LoginError)
	}
	//生成新token
	newToken, err := helper.GenToken(l.svcCtx.Config.JwtSecret, adminId, map[string]string{
		"RoleId": claims.Data["RoleId"],
	}, time.Duration(cKey.AdminTokenTtl))
	if err != nil {
		return nil, errx.NewMessageError(err.Error())
	}
	//缓存token
	tokenKey := cKey.AdminTokenKey + adminId
	l.svcCtx.Cache.Setex(tokenKey, token, cKey.AdminTokenTtl)
	//生成refreshToken
	str := adminId + newToken + time.Now().String() + l.svcCtx.Config.JwtSecret
	newRefreshToken := helper.Hash256(str)
	//缓存token 用来刷新的
	l.svcCtx.Cache.Setex(refreshTokenKey, refreshToken, cKey.AdminRefreshTokenTtl)
	return &types.LoginRes{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}, nil
}
