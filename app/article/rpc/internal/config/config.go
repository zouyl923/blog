package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Prefix   string
		Charset  string
	}
	KqViewNum struct {
		Brokers   []string
		Group     string
		Topic     string
		Offset    string
		Consumers int
	}
	UUidRpc zrpc.RpcClientConf
}
