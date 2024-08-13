package model

type RotationInput struct {
	PicUrl string `json:"pic_url"`
	Link   string `json:"link"`
	Sort   uint   `json:"sort"`
}

type RotationOutput struct {
	Id uint `json:"id"`
}
