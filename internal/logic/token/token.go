package token

import (
	"context"
	"gf-shop/internal/model"
	"gf-shop/internal/service"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sToken struct{}

func New() *sToken {
	return new(sToken)
}

var authService *jwt.GfJWTMiddleware

func (s *sToken) GfJWTMiddleware() *jwt.GfJWTMiddleware {
	return authService
}

func (s *sToken) Auth(r *ghttp.Request) {
	authService.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {

	service.RegisterToken(New())

	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "GfShop",
		Key:             []byte("test"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "UserId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

func Authenticator(ctx context.Context) (interface{}, error) {

	var (
		r  = g.RequestFromCtx(ctx)
		in model.AdminLoginInput
	)

	if err := r.Parse(&in); err != nil {
		return "", err
	}

	if val, err := service.Auth().LoginByNamePassword(ctx, in); err == nil {
		return val, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}
