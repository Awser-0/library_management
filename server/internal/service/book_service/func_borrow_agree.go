package book_service

import (
	"library/internal/model/entity"
	"library/internal/utils/result"
	"time"
)

func (s *BookService) BorrowAgree(recordId int64, desc string) (bool, *result.Result) {
	var record = s.borrowRecordDao.SelectRecord(recordId)
	if record == nil {
		return false, &result.BorrowRecordNotFound
	}
	if record.State != entity.BorrowRecordStateApply {
		return false, nil
	}
	record.State = entity.BorrowRecordStateBorrow
	record.ReplyDesc = desc
	var now = time.Now()
	record.ReplyTime = &now
	var ok = s.borrowRecordDao.UpdateRecord(record.Id, *record)
	return ok, nil
}
