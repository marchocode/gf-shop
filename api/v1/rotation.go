package v1

import "github.com/gogf/gf/v2/frame/g"

type RotationInfoReq struct {
	g.Meta `path:"/backend/rotation" method:"POST" summary:"Rotation Save"`
	PicUrl string `json:"pic_url" v:"required"`
	Link   string `json:"link" v:"required"`
	Sort   uint   `json:"sort" v:"required"`
}

type RotationInfoRes struct {
	Id uint `json:"id"`
}

type RotationInfoDelReq struct {
	g.Meta `path:"/backend/rotation/{Id}" method:"delete" summary:"delete rotation"`
	Id     uint `json:"id" in:"path" dc:"id"`
}

type RotationInfoDelRes struct {
}

type RotationInfoUpdateReq struct {
	g.Meta `path:"/backend/rotation/{Id}" method:"put" summary:"update rotation"`
	Id     uint   `json:"id" in:"path" dc:"id"`
	PicUrl string `json:"pic_url" v:"required"`
	Link   string `json:"link" v:"required"`
	Sort   uint   `json:"sort" v:"required"`
}

type RotationInfoUpdateRes struct{}

type RotationInfoListReq struct {
	g.Meta `path:"/backend/rotation" method:"get" summary:"get list"`
	Sort   int `json:"sort" in:"query"`
	CommonPageReq
}

type RotationInfoListRes struct {
	List interface{}
}

