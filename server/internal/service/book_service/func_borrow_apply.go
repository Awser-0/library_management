package book_service

import (
	"library/internal/model/entity"
	"library/internal/utils/result"
)

func (s *BookService) BorrowApply(bookUUID, userID int64, desc string) (bool, *result.Result) {
	var book = s.bookDao.SelectBook(bookUUID)
	if book == nil {
		return false, &result.BookNotFound
	}
	var ok = s.borrowRecordDao.InsertRecord(entity.BorrowRecord{
		BookUUID:  bookUUID,
		UserId:    userID,
		State:     entity.BorrowRecordStateApply,
		ApplyDesc: desc,
	})
	return ok, nil
}
