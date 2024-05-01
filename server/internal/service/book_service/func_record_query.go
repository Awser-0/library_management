package book_service

import (
	"library/internal/model/do"
	"library/internal/model/entity"
)

func (s *BookService) RecordQuery(page do.QueryPage) do.PageData[do.BorrowRecordDetail] {
	return s.borrowRecordDao.SelectRecordsInDetail(entity.BorrowRecord{}, page)
}

// 通过书籍UUID查询借阅记录
func (s *BookService) RecordQueryByBookUUID(bookUUID int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail] {
	return s.borrowRecordDao.SelectRecordsInDetail(entity.BorrowRecord{
		BookUUID: bookUUID,
	}, page)
}

// 通过用户id查询借阅记录
func (s *BookService) RecordQueryByUserID(userID int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail] {
	return s.borrowRecordDao.SelectRecordsInDetail(entity.BorrowRecord{
		UserId: userID,
	}, page)
}

// 通过书籍UUID查询自己的借阅记录，UUID可为空
func (s *BookService) RecordQuerySelfByBookUUID(userID int64, bookUUID *int64, page do.QueryPage) do.PageData[do.BorrowRecordDetail] {
	var query = entity.BorrowRecord{
		UserId: userID,
	}
	if bookUUID != nil {
		query.BookUUID = *bookUUID
	}
	return s.borrowRecordDao.SelectRecordsInDetail(query, page)
}
