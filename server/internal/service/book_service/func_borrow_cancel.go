package book_service

import (
	"library/internal/model/entity"
	"library/internal/utils/result"
	"time"
)

func (s *BookService) BorrowCancel(recordId int64, userId int64) (bool, *result.Result) {
	var record = s.borrowRecordDao.SelectRecord(recordId)
	if record == nil {
		return false, &result.BorrowRecordNotFound
	}
	if record.State != entity.BorrowRecordStateApply {
		return false, &result.BorrowRecordNotFound
	}
	record.State = entity.BorrowRecordStateCancel
	var now = time.Now()
	record.ReturnTime = &now
	var ok = s.borrowRecordDao.UpdateRecord(record.Id, *record)
	return ok, nil
}
