package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/service"
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

func (c *cAuth) RefreshToken(ctx context.Context, req *v1.AdminRefreshTokenReq) (res *v1.AdminRefreshTokenRes, err error) {
	res = &v1.AdminRefreshTokenRes{}
	res.Token, res.Expire = service.Token().GfJWTMiddleware().RefreshHandler(ctx)
	return
}

func (c *cAuth) Logout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error) {
	service.Token().GfJWTMiddleware().LogoutHandler(ctx)
	return
}
