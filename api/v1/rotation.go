package v1

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/backend/rotation" method:"POST" summary:"Rotation Save"`
	PicUrl string `json:"pic_url" v:"required"`
	Sort   uint   `json:"sort" v:"required"`
}

type RotationRes struct {
	Id uint `json:"id"`
}
