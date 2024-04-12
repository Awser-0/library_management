package service

import (
	"library/internal/model/entity"
	"library/internal/service/book_service"
	"library/internal/utils/result"
	"time"
)

type IBookService interface {
	BookQuery(queryString string) []entity.Book
	BookAdd(title, author, cover, introduction string, publishTime *time.Time, total int) bool
	BorrowAgree(recordId int64, desc string) (bool, *result.Result)
	BorrowApply(bookUUID, userID int64, desc string) (bool, *result.Result)
	BorrowReject(recordId int64, desc string) (bool, *result.Result)
	BorrowReturn(recordId int64) (bool, *result.Result)
	RecordQuery() []entity.BorrowRecord
	RecordQueryByBookUUID(bookUUID int64) []entity.BorrowRecord
	RecordQueryByUserID(userID int64) []entity.BorrowRecord
}

func NewBookService() IBookService {
	return book_service.New()
}
