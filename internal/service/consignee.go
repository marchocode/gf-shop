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
	IConsignee interface {
		Craete(ctx context.Context, input *model.ConsigneeCreateInput) (out *model.ConsigneeCreateOutput, err error)
	}
)

var (
	localConsignee IConsignee
)

func Consignee() IConsignee {
	if localConsignee == nil {
		panic("implement not found for interface IConsignee, forgot register?")
	}
	return localConsignee
}

func RegisterConsignee(i IConsignee) {
	localConsignee = i
}
