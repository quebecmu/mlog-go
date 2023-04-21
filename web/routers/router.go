package routers

import "github.com/cloudwego/hertz/pkg/app/server"

// Init 初始化所有路由组
func Init(h *server.Hertz) {

	TestV1Router(h)
	TestV2Router(h)
}