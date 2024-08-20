// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleWare interface {
		Ctx(r *ghttp.Request)
	}
)

var (
	localMiddleWare IMiddleWare
)

func MiddleWare() IMiddleWare {
	if localMiddleWare == nil {
		panic("implement not found for interface IMiddleWare, forgot register?")
	}
	return localMiddleWare
}

func RegisterMiddleWare(i IMiddleWare) {
	localMiddleWare = i
}
