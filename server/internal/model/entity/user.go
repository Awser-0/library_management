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
	Name       string     `orm:"profile_name" json:"profile_name"`
	Sex        string     `orm:"profile_sex" json:"profile_sex"`
	Phone      string     `orm:"profile_phone" json:"profile_phone"`
	Birth      *time.Time `orm:"profile_birth" json:"profile_birth"`
	IsAdmin    bool       `orm:"admin_flag" json:"admin_flag"`
	UpdateTime time.Time  `orm:"update_time" json:"update_time"`
	CreateTime time.Time  `orm:"create_time" json:"create_time"`
}
