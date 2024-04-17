package entity

import (
	"time"
)

// state: 申请借阅状态[ apply | cancel | reject | borrow | return ]
const (
	BorrowRecordStateApply  = "apply"
	BorrowRecordStateCancel = "cancel"
	BorrowRecordStateReject = "reject"
	BorrowRecordStateBorrow = "borrow"
	BorrowRecordStateReturn = "return"
)

type BorrowRecord struct {
	Id         int64      `orm:"id" json:"id"`
	UserId     int64      `orm:"user_id" json:"user_id"`
	BookUUID   int64      `orm:"book_uuid" json:"book_uuid"`
	State      string     `orm:"state" json:"state"`
	ReplyTime  *time.Time `orm:"reply_time" json:"reply_time"`
	ReturnTime *time.Time `orm:"return_time" json:"return_time"`
	ApplyDesc  string     `orm:"apply_description" json:"apply_description"`
	ReplyDesc  string     `orm:"reply_description" json:"reply_description"`
	UpdateTime time.Time  `orm:"update_time" json:"update_time"`
	CreateTime time.Time  `orm:"create_time" json:"create_time"`
}
