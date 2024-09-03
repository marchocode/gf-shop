package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

type AdminInfoCreateUpdateBase struct {
	Name     string
	Password string
	RoleIds  string
	UserSalt string
	IsAdmin  int
}

type AdminInfoCreateInput struct {
	AdminInfoCreateUpdateBase
}

type AdminInfoCreateOutput struct {
	Id int `json:"id"`
}

type AdminInfoWithRolesOutput struct {
	gmeta.Meta `orm:"table:admin_info"`
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	IsAdmin    int            `json:"isAdmin"`
	RoleIds    string         `json:"roleIds"`
	Roles      *AdminInfoRole `orm:"with:id=role_ids"`
}

type AdminInfoRole struct {
	gmeta.Meta `orm:"table:role_info"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
}
