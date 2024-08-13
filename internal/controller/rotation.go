package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

type cRotation struct{}

var Rotation = cRotation{}

func (a *cRotation) Create(ctx context.Context, req *v1.RotationReq) (res *v1.RotationRes, err error) {

	err = service.Rotation().Create(ctx, model.RotationInput{
		PicUrl: req.PicUrl,
		Link:   req.Link,
		Sort:   req.Sort,
	})

	return
}
