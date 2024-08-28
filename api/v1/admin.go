package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminInfoCreateReq struct {
	g.Meta   `path:"/backend/admin" method:"POST" summary:"Position Save"`
	Name     string `json:"name"         dc:"用户名"`
	Password string `json:"password"     dc:"密码"`
	RoleIds  string `json:"roleIds"      dc:"角色ids"`
	IsAdmin  int    `json:"isAdmin"      dc:"是否超级管理员"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/backend/admin/info" method:"get"`
}

type AdminGetInfoRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}

type AdminInfoCreateRes struct {
	Id int `json:"id"`
}
