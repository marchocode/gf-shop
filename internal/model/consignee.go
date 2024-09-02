package model

type ConsigneeCreateInput struct {
	UserId    int    `json:"userId"`
	IsDefault int    `json:"isDefault"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Town      string `json:"town"`
	Street    string `json:"street"`
	Detail    string `json:"detail"`
}

type ConsigneeCreateOutput struct {
	Id int `json:"id"`
}
