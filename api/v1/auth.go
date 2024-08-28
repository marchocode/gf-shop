package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta   `path:"/backend/auth/login" method:"POST" summary:"Admin Login"`
	Name     string `json:"name"         dc:"用户名"`
	Password string `json:"password"     dc:"密码"`
}

type AdminLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AdminLogoutReq struct {
	g.Meta `path:"/backend/auth/logout" method:"post"`
}

type AdminLogoutRes struct {
}

type AdminRefreshTokenReq struct {
	g.Meta `path:"/backend/auth/refresh_token" method:"post"`
}

type AdminRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
