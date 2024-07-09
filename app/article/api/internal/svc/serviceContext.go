package svc

import (
	"blog/app/article/api/internal/config"
	"blog/app/article/api/internal/middleware"
	"blog/app/article/rpc/client/adminservice"
	"blog/app/article/rpc/client/commonservice"
	"blog/app/article/rpc/client/userservice"
	"blog/app/verify/rpc/client/verifyservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	CorsMiddleware      rest.Middleware
	UserAuthMiddleware  rest.Middleware
	AdminAuthMiddleware rest.Middleware

	ArticleCommonRpc commonservice.CommonService
	ArticleUserRpc   userservice.UserService
	ArticleAdminRpc  adminservice.AdminService
}

func NewServiceContext(c config.Config) *ServiceContext {
	articleCommonRpc := commonservice.NewCommonService(zrpc.MustNewClient(c.ArticleRpc))
	articleUserRpc := userservice.NewUserService(zrpc.MustNewClient(c.ArticleRpc))
	articleAdminRpc := adminservice.NewAdminService(zrpc.MustNewClient(c.ArticleRpc))
	verifyRpc := verifyservice.NewVerifyService(zrpc.MustNewClient(c.VerifyRpc))
	return &ServiceContext{
		Config:             c,
		CorsMiddleware:     middleware.NewCorsMiddleware().Handle,
		UserAuthMiddleware: middleware.NewUserAuthMiddleware(c, verifyRpc).Handle,
		ArticleCommonRpc:   articleCommonRpc,
		ArticleUserRpc:     articleUserRpc,
		ArticleAdminRpc:    articleAdminRpc,
	}
}
