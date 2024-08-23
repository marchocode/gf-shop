package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/service"
)

type cAuth struct{}

var Auth = cAdmin{}

func (a *cAdmin) Login(ctx context.Context, req *v1.AdminLoginReq) (*v1.AdminLoginRes, error) {

	token, expire := service.Token().GfJWTMiddleware().LoginHandler(ctx)

	return &v1.AdminLoginRes{
		Token:  token,
		Expire: expire,
	}, nil
}
