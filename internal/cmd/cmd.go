package cmd

import (
	"context"
	"gf-shop/internal/controller"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)

				group.Bind(controller.Auth)
			})

			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Token().Auth,
				)

				group.Bind(
					controller.Admin,
					controller.Rotation,
					controller.Position,
				)
			})

			s.Run()
			return nil
		},
	}
)
