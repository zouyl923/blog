package middleware

import (
	"blog/app/admin/api/internal/config"
	"blog/common/helper"
	"blog/common/response"
	"blog/common/response/errx"
	"blog/common/static/cKey"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type AuthMiddleware struct {
	Config config.Config
	Cache  *redis.Redis
}

func NewAuthMiddleware(c config.Config, redis *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		Config: c,
		Cache:  redis,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		//解析token
		claims, err := helper.ParseToken(token, m.Config.JwtSecret)
		if err != nil {
			resp := response.Response{
				Code:    errx.LoginError,
				Message: errx.GetCnMessage(errx.LoginError),
			}
			str, _ := json.Marshal(resp)
			w.Write(str)
			return
		}
		adminId := claims.Key
		//验证token是否有效
		tokenKey := cKey.AdminTokenKey + adminId
		cToken, _ := m.Cache.Get(tokenKey)
		//验证token是否有效
		refreshTokenKey := cKey.AdminRefreshTokenKey + adminId
		cRefreshToken, _ := m.Cache.Get(refreshTokenKey)
		if len(cToken) == 0 || len(cRefreshToken) == 0 || cToken != token {
			resp := response.Response{
				Code:    errx.LoginExpire,
				Message: errx.GetCnMessage(errx.LoginExpire),
			}
			str, _ := json.Marshal(resp)
			w.Write(str)
			return
		}
		//追加参数
		ctx := r.Context()
		ctx = context.WithValue(ctx, "adminId", adminId)
		roleId := claims.Data["RoleId"]
		ctx = context.WithValue(ctx, "roleId", roleId)
		newR := r.WithContext(ctx)
		next(w, newR)
	}
}
