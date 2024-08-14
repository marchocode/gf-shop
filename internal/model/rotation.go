package model

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
