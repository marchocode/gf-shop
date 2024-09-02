package v1

import "github.com/gogf/gf/v2/frame/g"

type OrderTodaysNumReq struct {
	g.Meta `path:"/backend/order/nums/today" method:"GET" summary:"Orders num of today"`
}

type OrderTodaysNumRes struct {
	Nums int    `json:"nums" dc:"nums"`
	Date string `json:"date" dc:"date"`
}
