package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xxandjg/mlog-go/logger"
	"github.com/xxandjg/mlog-go/web/config"
	"github.com/xxandjg/mlog-go/web/routers"
)

func main() {

	// 加载配置文件
	if err := config.Init("dev"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		panic(err)
		return
	}
	if err := logger.Init(config.Conf.Log, "dev"); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		panic(err)
		return
	}

	h := server.Default(
		server.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", config.Conf.Port)),
	)

	routers.Init(h)

	h.Spin()

}
