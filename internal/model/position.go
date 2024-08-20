package model

type PositionCreateUpdateBase struct {
	PicUrl    string `json:"pic_url"`
	GoodsName string `json:"goods_name"`
	Link      string `json:"link"`
	Sort      int    `json:"sort"`
	GoodsId   int    `json:"goods_id"`
}

type PositionInfoCreateInput struct {
	PositionCreateUpdateBase
}

type PositionInfoUpdateInput struct {
	PositionCreateUpdateBase
	Id int
}

type PositionInfoCreateOutput struct {
	Id int
}
