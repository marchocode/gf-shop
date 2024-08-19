package v1

type CommonPageReq struct {
	Page int `json:"page" d:"1"  v:"required"`
	Size int `json:"size" d:"10" v:"required"`
}

type CommonPageRes struct {
	Total int `json:"total" dc:"total"`
}
