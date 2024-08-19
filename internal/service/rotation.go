// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-shop/internal/model"
)

type (
	IRotationInfo interface {
		Create(ctx context.Context, in model.RotationInfoCreateInput) error
		Delete(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, in model.RotationInfoUpdateInput) (err error)
		List(ctx context.Context, in model.RotationInfoListInput) (out *model.RotationInfoListOutput, err error)
	}
)

var (
	localRotationInfo IRotationInfo
)

func RotationInfo() IRotationInfo {
	if localRotationInfo == nil {
		panic("implement not found for interface IRotationInfo, forgot register?")
	}
	return localRotationInfo
}

func RegisterRotationInfo(i IRotationInfo) {
	localRotationInfo = i
}
