package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

type cAuth struct{}

var Auth = cAuth{}

func (a *cAuth) Login(ctx context.Context, req *v1.AdminLoginReq) (*v1.AdminLoginRes, error) {

	token, expire := service.Token().GfJWTMiddleware().LoginHandler(ctx)

	return &v1.AdminLoginRes{
		Token:  token,
		Expire: expire,
	}, nil
}

func Login(r *ghttp.Request) (string, interface{}) {

	input := model.AdminLoginInput{
		Name:     r.Get("name").String(),
		Password: r.Get("password").String(),
	}

	gutil.Dump(input)

	m, err := service.Auth().LoginByNamePassword(r.GetCtx(), input)

	if err != nil {
		r.Response.WriteJson(gtoken.Fail(err.Error()))
		r.ExitAll()
	}

	return gconv.String(m["UserId"]), m
}

func (c *cAuth) RefreshToken(ctx context.Context, req *v1.AdminRefreshTokenReq) (res *v1.AdminRefreshTokenRes, err error) {
	res = &v1.AdminRefreshTokenRes{}
	res.Token, res.Expire = service.Token().GfJWTMiddleware().RefreshHandler(ctx)
	return
}

func (c *cAuth) Logout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error) {
	service.Token().GfJWTMiddleware().LogoutHandler(ctx)
	return
}
