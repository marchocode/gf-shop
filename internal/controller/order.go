package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/service"
)

type cOrder struct{}

var Order = &cOrder{}

func (c *cOrder) TodaysOrders(ctx context.Context, req *v1.OrderTodaysNumReq) (*v1.OrderTodaysNumRes, error) {

	out, err := service.Order().TodaysOrders(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.OrderTodaysNumRes{
		Nums: out.Nums,
		Date: out.Date,
	}, nil

}
