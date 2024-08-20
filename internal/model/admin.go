package model

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
