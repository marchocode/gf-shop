package controller

import (
	"context"
	v1 "gf-shop/api/v1"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

type cUpload struct{}

var Upload = cUpload{}

func (c *cUpload) Upload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {

	upload := g.Cfg().MustGet(ctx, "upload.path").String()
	filename, err := req.File.Save(upload, true)

	if err != nil {
		return
	}

	out, err := service.Qiniu().Upload(ctx, &model.FileUploadInput{
		TempPath: gfile.Join(upload, filename),
		FileName: req.File.Filename,
	})

	if err != nil {
		return
	}

	res = &v1.FileUploadRes{
		Name: filename,
		Url:  out.Key,
	}

	return
}
