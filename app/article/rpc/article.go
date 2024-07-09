package main

import (
	"flag"
	"fmt"

	"blog/app/article/rpc/internal/config"
	adminserviceServer "blog/app/article/rpc/internal/server/adminservice"
	commonserviceServer "blog/app/article/rpc/internal/server/commonservice"
	userserviceServer "blog/app/article/rpc/internal/server/userservice"
	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterCommonServiceServer(grpcServer, commonserviceServer.NewCommonServiceServer(ctx))
		rpc.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		rpc.RegisterAdminServiceServer(grpcServer, adminserviceServer.NewAdminServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
