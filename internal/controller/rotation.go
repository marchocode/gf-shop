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

	err = service.RotationInfo().Create(ctx, model.RotationInfoCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})

	return
}

// Delete Rotation by id
func (a *cRotation) Delete(ctx context.Context, req *v1.RotationInfoDelReq) (res *v1.RotationInfoDelRes, err error) {
	err = service.RotationInfo().Delete(ctx, req.Id)
	return
}

// Update Rotation
func (a *cRotation) Update(ctx context.Context, req *v1.RotationInfoUpdateReq) (res *v1.RotationInfoUpdateRes, err error) {

	err = service.RotationInfo().Update(ctx, model.RotationInfoUpdateInput{
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})

	return
}

// Get List
func (a *cRotation) List(ctx context.Context, req *v1.RotationInfoListReq) (res *v1.RotationInfoListRes, err error) {

	list, err := service.RotationInfo().List(ctx, model.RotationInfoListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &v1.RotationInfoListRes{
		List: list.List,
		CommonPageRes: v1.CommonPageRes{
			Total: list.Total,
		},
	}, nil

}
