package book_service

import "library/internal/model/entity"

func (s *BookService) RecordQuery() []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{})
}

func (s *BookService) RecordQueryByBookUUID(bookUUID int64) []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{
		BookUUID: bookUUID,
	})
}

func (s *BookService) RecordQueryByUserID(userID int64) []entity.BorrowRecord {
	return s.borrowRecordDao.SelectRecordsRestricted(entity.BorrowRecord{
		UserId: userID,
	})
}
