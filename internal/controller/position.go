package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

type cPosition struct{}

var Position = cPosition{}

func (a *cPosition) Create(ctx context.Context, req *v1.PositionInfoCreateReq) (*v1.PositionInfoCreateRes, error) {

	output, err := service.Position().Create(ctx, model.PositionInfoCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			GoodsName: req.GoodsName,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
		},
	})

	if err != nil {
		return nil, err
	}

	return &v1.PositionInfoCreateRes{
		Id: output.Id,
	}, nil
}

// Update Position
func (a *cPosition) Update(ctx context.Context, req *v1.PositionInfoUpdateReq) (res *v1.PositionInfoUpdateRes, err error) {

	err = service.Position().Update(ctx, model.PositionInfoUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			GoodsName: req.GoodsName,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
		},
	})

	return
}
