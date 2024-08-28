package controller

import (
	"context"
	"fmt"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/util/gconv"
)

type cAdmin struct{}

var Admin = cAdmin{}

func (a *cAdmin) Create(ctx context.Context, req *v1.AdminInfoCreateReq) (*v1.AdminInfoCreateRes, error) {

	output, err := service.Admin().Create(ctx, model.AdminInfoCreateInput{
		AdminInfoCreateUpdateBase: model.AdminInfoCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})

	if err != nil {
		return nil, err
	}

	return &v1.AdminInfoCreateRes{
		Id: output.Id,
	}, nil
}

func (c *cAdmin) Info(ctx context.Context, req *v1.AdminGetInfoReq) (*v1.AdminGetInfoRes, error) {

	auth := service.Token().GfJWTMiddleware()

	res := &v1.AdminGetInfoRes{
		Id:          gconv.Int(auth.GetIdentity(ctx)),
		IdentityKey: auth.IdentityKey,
		Payload:     auth.GetPayload(ctx),
	}

	fmt.Println(res)

	return res, nil
}
