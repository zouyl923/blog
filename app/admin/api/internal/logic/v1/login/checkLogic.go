package login

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/common/static/cKey"
	"context"
	"net/http"
	"time"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证是否登录成功，并刷新token
func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(r *http.Request) (resp *types.LoginCheckRes, err error) {
	adminId := l.ctx.Value("adminId").(string)
	roleId := l.ctx.Value("roleId").(string)
	token := r.Header.Get("token")
	refreshToken := r.Header.Get("refresh-token")

	refreshTokenKey := cKey.AdminRefreshTokenKey + adminId
	cRefreshToken, _ := l.svcCtx.Cache.Get(refreshTokenKey)
	if cRefreshToken == "" || cRefreshToken != refreshToken {
		return nil, errx.NewCodeError(errx.LoginError)
	}
	//缓存token
	tokenKey := cKey.AdminTokenKey + adminId
	expire, _ := l.svcCtx.Cache.Ttl(tokenKey)

	newToken := token
	newRefreshToken := refreshToken
	//如果Token快到期了(提前10分钟替换)，或者token有效期出现异常，立马更新token
	if expire < 10*60 || expire > cKey.AdminTokenTtl {
		//生成新token
		newToken, err = helper.GenToken(l.svcCtx.Config.JwtSecret, adminId, map[string]string{
			"RoleId": roleId,
		}, time.Duration(cKey.AdminTokenTtl))
		if err != nil {
			return nil, errx.NewMessageError(err.Error())
		}
		l.svcCtx.Cache.Setex(tokenKey, newToken, cKey.AdminTokenTtl)
		//生成refreshToken
		str := adminId + newToken + time.Now().String() + l.svcCtx.Config.JwtSecret
		newRefreshToken = helper.Hash256(str)
		//缓存token 用来刷新的
		l.svcCtx.Cache.Setex(refreshTokenKey, newRefreshToken, cKey.AdminRefreshTokenTtl)
	}
	return &types.LoginCheckRes{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}, nil
}
