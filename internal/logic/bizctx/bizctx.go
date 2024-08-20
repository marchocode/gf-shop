package bizctx

import (
	"context"
	"gf-shop/internal/consts"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sBizCtx struct{}

func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return new(sBizCtx)
}

// 初始化自定义上下文到 request.
func (s *sBizCtx) Init(r *ghttp.Request, ctx model.Context) {
	r.SetCtxVar(consts.ContextKey, ctx)
}

// 从上下文获得自定义上下文
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)

	if value == nil {
		return nil
	}

	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
