package v1

import "github.com/gogf/gf/v2/frame/g"

type PositionInfoCreateReq struct {
	g.Meta    `path:"/backend/position" method:"POST" summary:"Position Save"`
	PicUrl    string `json:"picUrl" v:"required"      dc:"图片链接"`
	GoodsName string `json:"goodsName" v:"required" dc:"商品名称"`
	Link      string `json:"link"     v:"required"        dc:"跳转链接"`
	Sort      int    `json:"sort"    v:"required"        dc:"排序"`
	GoodsId   int    `json:"goods_id" v:"required"`
}

type PositionInfoUpdateReq struct {
	g.Meta    `path:"/backend/position" method:"put" summary:"Position Update"`
	Id        int    `json:"id" v:"required"`
	PicUrl    string `json:"picUrl" v:"required"      dc:"图片链接"`
	GoodsName string `json:"goodsName" v:"required" dc:"商品名称"`
	Link      string `json:"link"     v:"required"        dc:"跳转链接"`
	Sort      int    `json:"sort"    v:"required"        dc:"排序"`
	GoodsId   int    `json:"goods_id" v:"required"`
}

type PositionInfoUpdateRes struct {
	Id int `json:"id"`
}

type PositionInfoCreateRes struct {
	Id int `json:"id"`
}
