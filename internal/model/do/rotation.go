// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Rotation is the golang structure of table rotation for DAO operations like Where/Data.
type Rotation struct {
	g.Meta   `orm:"table:rotation, do:true"`
	Id       interface{} // id
	PicUrl   interface{} // url
	Sort     interface{} // sort
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
