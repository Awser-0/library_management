package entity

import "time"

const (
	Book_TableNA = 1
)

type Book struct {
	UUID         int64      `orm:"uuid" json:"uuid"`
	Title        string     `orm:"title" json:"title"`
	Author       string     `orm:"author" json:"author"`
	Cover        string     `orm:"cover" json:"cover"`
	Introduction string     `orm:"introduction" json:"introduction"`
	PublishTime  *time.Time `orm:"publish_time" json:"publish_time"`
	Total        int        `orm:"total" json:"total"`
	UpdateTime   time.Time  `orm:"update_time" json:"update_time"`
	CreateTime   time.Time  `orm:"create_time" json:"create_time"`
}
