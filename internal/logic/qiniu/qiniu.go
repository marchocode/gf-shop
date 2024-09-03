package qiniu

import (
	"context"
	"gf-shop/internal/model"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type sQiniu struct{}

func init() {
	service.RegisterQiniu(New())
}

func (s *sQiniu) Upload(ctx context.Context, in *model.FileUploadInput) (out *model.FileUploadOutput, err error) {

	accessKey := g.Cfg().MustGet(ctx, "qiniu.ak").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.sk").String()

	putPolicy := storage.PutPolicy{
		Scope: "marcho",
	}

	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err = formUploader.PutFile(context.Background(), &ret, upToken, in.FileName, in.TempPath, &putExtra)

	if err != nil {
		return
	}

	out = &model.FileUploadOutput{
		Key:  ret.Key,
		Hash: ret.Hash,
	}

	return
}

func New() *sQiniu {
	return new(sQiniu)
}
