package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type cAuth struct{}

var Auth = cAdmin{}

func (a *cAdmin) Login(ctx context.Context, req *v1.AdminLoginReq) (*v1.AdminLoginRes, error) {

	g.Dump(req)

	err := service.Auth().Login(ctx, model.AdminLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &v1.AdminLoginRes{
		User: service.BizCtx().Get(ctx).User,
	}, nil

}
