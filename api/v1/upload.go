package v1

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
)

type FileUploadReq struct {
	g.Meta `path:"/backend/file/upload" method:"post" mime:"multipart/form-data" tags:"工具" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url"  dc:"访问URL，可能只是URI"`
}
