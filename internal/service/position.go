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
	IPosition interface {
		Create(ctx context.Context, in model.PositionInfoCreateInput) (*model.PositionInfoCreateOutput, error)
	}
)

var (
	localPosition IPosition
)

func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}

func RegisterPosition(i IPosition) {
	localPosition = i
}
