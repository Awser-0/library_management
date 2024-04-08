package entity

import "time"

const (
	UserSexBoy  = "0"
	UserSexGirl = "1"
)

type User struct {
	Id         int64      `orm:"id"`
	Username   string     `orm:"username"`
	Password   string     `orm:"password"`
	Sex        string     `orm:"profile_sex"`
	Phone      string     `orm:"profile_phone"`
	Birth      *time.Time `orm:"profile_birth"`
	IsAdmin    bool       `orm:"admin_flag"`
	UpdateTime time.Time  `orm:"update_time"`
	CreateTime time.Time  `orm:"create_time"`
}
