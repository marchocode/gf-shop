package v1

import "github.com/gogf/gf/v2/frame/g"

type RotationInfoReq struct {
	g.Meta `path:"/backend/rotation" method:"POST" summary:"Rotation Save"`
	PicUrl string `json:"pic_url" v:"required"`
	Link   string `json:"link" v:"required"`
	Sort   uint   `json:"sort" v:"required"`
}

type RotationInfoDelReq struct {
	g.Meta `path:"/backend/rotation/{Id}" method:"delete" summary:"delete rotation" tags:"用户"`
	Id     uint `json:"id" in:"path" dc:"id"`
}

type RotationInfoDelRes struct {
}

type RotationInfoRes struct {
	Id uint `json:"id"`
}
