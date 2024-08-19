package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type RotationCreateUpdateBase struct {
	PicUrl string `json:"pic_url"`
	Link   string `json:"link"`
	Sort   uint   `json:"sort"`
}

type RotationInfoCreateInput struct {
	RotationCreateUpdateBase
}

type RotationInfoUpdateInput struct {
	RotationCreateUpdateBase
	Id uint `json:"id"`
}

type RotationInfoOutput struct {
	Id uint `json:"id"`
}

type RotationInfoListInput struct {
	Page int
	Size int
	Sort int
}

type RotationInfoListOutput struct {
	List  []RotationInfoListItem
	Total int
}

type RotationInfoListItem struct {
	Id        int         `json:"id"`
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      uint        `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}
