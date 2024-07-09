package svc

import (
	"blog/app/user/api/internal/config"
	"blog/app/user/api/internal/middleware"
	"blog/app/user/rpc/client/userservice"
	"blog/app/verify/rpc/client/verifyservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	CorsMiddleware     rest.Middleware
	UserAuthMiddleware rest.Middleware
	UserRpc            userservice.UserService
	VerifyRpc          verifyservice.VerifyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := userservice.NewUserService(zrpc.MustNewClient(c.UserRpc))
	verifyRpc := verifyservice.NewVerifyService(zrpc.MustNewClient(c.VerifyRpc))
	return &ServiceContext{
		Config:             c,
		CorsMiddleware:     middleware.NewCorsMiddleware().Handle,
		UserAuthMiddleware: middleware.NewUserAuthMiddleware(c, verifyRpc).Handle,
		UserRpc:            userRpc,
	}
}
