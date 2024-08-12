// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Rotation is the golang structure for table rotation.
type Rotation struct {
	Id       uint        `json:"id"       orm:"id"        description:"id"`
	PicUrl   string      `json:"picUrl"   orm:"pic_url"   description:"url"`
	Sort     uint        `json:"sort"     orm:"sort"      description:"sort"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" description:""`
}
