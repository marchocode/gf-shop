package cmd

import (
	"context"
	"gf-shop/internal/controller"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var gfToken *gtoken.GfToken

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			gfToken = &gtoken.GfToken{
				ServerName:      "gf shop",
				CacheMode:       2,
				LoginPath:       "/backend/auth/login",
				LoginBeforeFunc: controller.Login,
				LogoutPath:      "post:/backend/auth/logout",
			}

			// 不需要认证的请求
			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)

				group.Bind(
					controller.Order,
					controller.Consignee,
					controller.Upload,
				)

				group.ALL("/hello", func(r *ghttp.Request) {
					r.Response.WriteJson(gtoken.Succ("hello"))
				})

			})

			// 需要认证的请求
			s.Group("/", func(group *ghttp.RouterGroup) {

				err := gfToken.Middleware(ctx, group)

				if err != nil {
					panic(err)
				}

				group.Middleware(func(r *ghttp.Request) {

					// inject current user.
					resp := gfToken.GetTokenData(r)
					r.SetCtxVar("User", resp.Data)

					r.Middleware.Next()
				})

				group.Middleware(ghttp.MiddlewareHandlerResponse)

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
