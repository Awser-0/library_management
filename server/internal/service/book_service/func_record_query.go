package book_service

import "library/internal/model/entity"

func (s *BookService) RecordQuery() []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{})
}

// 通过书籍UUID查询借阅记录
func (s *BookService) RecordQueryByBookUUID(bookUUID int64) []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{
		BookUUID: bookUUID,
	})
}

// 通过用户id查询借阅记录
func (s *BookService) RecordQueryByUserID(userID int64) []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{
		UserId: userID,
	})
}

// 通过书籍UUID查询自己的借阅记录，UUID可为空
func (s *BookService) RecordQuerySelfByBookUUID(userID int64, bookUUID *int64) []entity.BorrowRecord {
	var query = entity.BorrowRecord{
		UserId: userID,
	}
	if bookUUID != nil {
		query.BookUUID = *bookUUID
	}
	return s.borrowRecordDao.SelectRecordsRestricted(query)
}
