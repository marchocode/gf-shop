package v1

import "github.com/gogf/gf/v2/frame/g"

type ConsigneeCreateUpdateBase struct {
	UserId    int    `json:"userId"         d:""`
	IsDefault int    `json:"isDefault"   d:"默认地址1  非默认0"`
	Name      string `json:"name"              d:"姓名"`
	Phone     string `json:"phone"           d:"手机号"`
	Province  string `json:"province"     d:"省份"`
	City      string `json:"city"             d:"城市"`
	Town      string `json:"town"              d:"县区"`
	Street    string `json:"street"         d:"街道乡镇"`
	Detail    string `json:"detail"         d:"地址详情"`
}

type ConsigneeCreateReq struct {
	g.Meta `path:"/backend/consignee" method:"POST"`
	ConsigneeCreateUpdateBase
}

type ConsigneeCreateRes struct {
	Id int `json:"id" d:"id"`
}
