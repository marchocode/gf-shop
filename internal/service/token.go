// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IToken interface {
		GfJWTMiddleware() *jwt.GfJWTMiddleware
		Auth(r *ghttp.Request)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
