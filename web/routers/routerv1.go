package routers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestV1Router(h *server.Hertz) {
	v1 := h.Group("/v1")
	v1.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "getv1")
	})
	v1.POST("/post", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "postv1")
	})
}