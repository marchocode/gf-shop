// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-shop/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IBizCtx interface {
		// 初始化自定义上下文到 request.
		Init(r *ghttp.Request, ctx model.Context)
		// 从上下文获得自定义上下文
		Get(ctx context.Context) *model.Context
		SetUser(ctx context.Context, ctxUser *model.ContextUser)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
