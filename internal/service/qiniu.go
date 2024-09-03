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
	IQiniu interface {
		Upload(ctx context.Context, in *model.FileUploadInput) (out *model.FileUploadOutput, err error)
	}
)

var (
	localQiniu IQiniu
)

func Qiniu() IQiniu {
	if localQiniu == nil {
		panic("implement not found for interface IQiniu, forgot register?")
	}
	return localQiniu
}

func RegisterQiniu(i IQiniu) {
	localQiniu = i
}
