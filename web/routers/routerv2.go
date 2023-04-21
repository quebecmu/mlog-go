package routers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestV2Router(h *server.Hertz) {
	v2 := h.Group("/v2")
	v2.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "getv2")
	})
	v2.POST("/post", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "postv2")
	})
}
