package session

import (
	"context"
	"gf-shop/internal/model/entity"
	"gf-shop/internal/service"

	"github.com/gogf/gf/v2/util/gutil"
)

const (
	sessionKeyUser = "sessionKeyUser"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return new(sSession)
}

func (s *sSession) SetUser(ctx context.Context, user *entity.AdminInfo) error {

	gutil.Dump("user", user)
	bizContext := service.BizCtx().Get(ctx)
	return bizContext.Session.Set(sessionKeyUser, user)
}

func (*sSession) GetUser(ctx context.Context) *entity.AdminInfo {

	// 获得自定义session
	customCtx := service.BizCtx().Get(ctx)

	if customCtx == nil {
		return &entity.AdminInfo{}
	}

	v, _ := customCtx.Session.Get(sessionKeyUser)

	if !v.IsNil() {
		var user *entity.AdminInfo
		_ = v.Struct(&user)
		return user
	}

	return &entity.AdminInfo{}
}
