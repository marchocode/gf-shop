package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

type cRotation struct{}

var Rotation = cRotation{}

// Create Rotation
func (a *cRotation) Create(ctx context.Context, req *v1.RotationInfoReq) (res *v1.RotationInfoRes, err error) {

	err = service.RotationInfo().Create(ctx, model.RotationInfoInput{
		PicUrl: req.PicUrl,
		Link:   req.Link,
		Sort:   req.Sort,
	})

	return
}

func (a *cRotation) Delete(ctx context.Context, req *v1.RotationInfoDelReq) (res *v1.RotationInfoDelRes, err error) {
	err = service.RotationInfo().Delete(ctx, req.Id)
	return
}
