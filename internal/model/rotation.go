package model

type RotationInfoInput struct {
	PicUrl string `json:"pic_url"`
	Link   string `json:"link"`
	Sort   uint   `json:"sort"`
}

type RotationInfoOutput struct {
	Id uint `json:"id"`
}
