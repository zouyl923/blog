package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	TencentCloud struct {
		SecretId  string
		SecretKey string
		Cos       struct {
			AppId  int
			Bucket string
			REGION string
			Domain string
		}
	}
	ArticleRpc zrpc.RpcClientConf
	VerifyRpc  zrpc.RpcClientConf
}
