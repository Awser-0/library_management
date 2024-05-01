package service

import (
	"library/internal/model/do"
	"library/internal/model/entity"
	"library/internal/service/book_service"
	"library/internal/utils/result"
	"time"
)

type IBookService interface {
	BookAdd(title, author, cover, introduction string, publishTime *time.Time, total int) bool
	BookQuery(queryString string) []entity.Book
	BookSelect(uuid int64) *entity.Book
	BookUpdate(book entity.Book) bool
	BorrowApply(bookUUID, userID int64, desc string) (bool, *result.Result)
	BorrowCancel(recordId int64, userId int64) (bool, *result.Result)
	BorrowAgree(recordId int64, desc string) (bool, *result.Result)
	BorrowReject(recordId int64, desc string) (bool, *result.Result)
	BorrowReturn(recordId int64) (bool, *result.Result)
	RecordQuery(page do.QueryPage) do.PageData[do.BorrowRecordDetail]
	RecordQueryByBookUUID(bookUUID int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail]
	RecordQueryByUserID(userID int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail]
	RecordQuerySelfByBookUUID(userID int64, bookUUID *int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail]
}

func NewBookService() IBookService {
	return book_service.New()
}
