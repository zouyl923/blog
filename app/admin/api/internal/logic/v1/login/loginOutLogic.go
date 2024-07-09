package login

import (
	"blog/app/admin/api/internal/svc"
	"blog/common/static/cKey"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 退出登录
func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginOutLogic) LoginOut() error {
	adminId := l.ctx.Value("adminId").(string)
	tokenKey := cKey.AdminTokenKey + adminId
	refreshTokenKey := cKey.AdminRefreshTokenKey + adminId
	//删除缓存
	l.svcCtx.Cache.Del(tokenKey)
	l.svcCtx.Cache.Del(refreshTokenKey)
	return nil
}
