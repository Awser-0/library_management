package entity

import (
	"time"

	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	UserSexBoy  = "0"
	UserSexGirl = "1"
)

type User struct {
	gmeta.Meta `orm:"table:user"`
	Id         int64      `orm:"id" json:"id"`
	Username   string     `orm:"username" json:"username"`
	Password   string     `orm:"password" json:"-"`
	Nickname   string     `orm:"nickname" json:"nickname"`
	Sex        string     `orm:"sex" json:"sex"`
	Phone      string     `orm:"phone" json:"phone"`
	Birth      *time.Time `orm:"birth" json:"birth"`
	IsAdmin    bool       `orm:"is_admin" json:"is_admin"`
	UpdateTime time.Time  `orm:"update_time" json:"update_time"`
	CreateTime time.Time  `orm:"create_time" json:"create_time"`
}
