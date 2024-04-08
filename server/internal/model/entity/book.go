package entity

import "time"

const (
	Book_TableNA = 1
)

type Book struct {
	UUID        int64     `orm:"uuid"`
	Title       string    `orm:"title"`
	Author      string    `orm:"author"`
	Cover       string    `orm:"cover"`
	PublishTime time.Time `orm:"publish_time"`
	Total       int       `orm:"total"`
	UpdateTime  time.Time `orm:"update_time"`
	CreateTime  time.Time `orm:"create_time"`
}
