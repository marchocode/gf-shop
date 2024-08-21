package middleware

import (
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleWare struct{}

func init() {
	service.RegisterMiddleWare(New())
}

func New() *sMiddleWare {
	return new(sMiddleWare)
}

func (s *sMiddleWare) Ctx(r *ghttp.Request) {

	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}

	service.BizCtx().Init(r, customCtx)

	if user := service.Session().GetUser(r.Context()); user.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:      user.Id,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
		}
	}

	r.Assign("ContextKey", customCtx)

	r.Middleware.Next()
}
