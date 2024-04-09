package dao

import (
	borrowrecorddaoimpl "library/internal/dao/borrow_record_dao_impl"
	"library/internal/model/entity"
)

type IBorrowRecordDao interface {
	SelectRecord(id int64) *entity.BorrowRecord
	SelectRecordsRestricted(record entity.BorrowRecord) []entity.BorrowRecord
	InsertRecord(borrowRecord entity.BorrowRecord) bool
	UpdateRecord(id int64, book entity.BorrowRecord) bool
	DeleteRecord(id int64) bool
}

func NewBorrowRecordDao() IBorrowRecordDao {
	return borrowrecorddaoimpl.New()
}
