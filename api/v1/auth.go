package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta   `path:"/backend/login" method:"POST" summary:"Admin Login"`
	Name     string `json:"name"         dc:"用户名"`
	Password string `json:"password"     dc:"密码"`
}

type AdminLoginRes struct {
	User interface{} `json:"user"`
}
