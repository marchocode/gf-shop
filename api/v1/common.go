package v1

type CommonPageReq struct {
	Page uint `json:"page" d:"1"  v:"required"`
	Size uint `json:"size" d:"10" v:"required"`
}

type CommonPageRes struct {
	Total uint `dc:"total"`
}
