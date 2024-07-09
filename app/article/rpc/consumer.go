package main

import (
	"blog/app/article/rpc/internal/config"
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var configFile = flag.String("f", "etc/article.yaml", "the config file")
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	kqConsumer := kq.MustNewQueue(kq.KqConf{
		Brokers: c.KqViewNum.Brokers,
		Group:   c.KqViewNum.Group,
		Topic:   c.KqViewNum.Topic,
	}, kq.WithHandle(func(key, value string) error {
		fmt.Printf("=> %s\n", value)
		return nil
	}))
	fmt.Printf("%+v", kqConsumer)
	kqConsumer.Start()
	defer kqConsumer.Stop()
}
