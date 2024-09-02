package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

type cConsignee struct{}

var Consignee = cConsignee{}

func (c *cConsignee) Create(ctx context.Context, req *v1.ConsigneeCreateReq) (res *v1.ConsigneeCreateRes, err error) {

	out, err := service.Consignee().Craete(ctx, &model.ConsigneeCreateInput{
		UserId:    req.UserId,
		IsDefault: req.IsDefault,
		Name:      req.Name,
		Phone:     req.Phone,
		Province:  req.Province,
		City:      req.City,
		Town:      req.Town,
		Street:    req.Street,
		Detail:    req.Detail,
	})

	res = new(v1.ConsigneeCreateRes)
	res.Id = out.Id

	return
}
