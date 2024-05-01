package borrowrecorddaoimpl

import (
	"library/internal/model/do"
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BorrowRecordDaoImpl) SelectRecord(id int64) *entity.BorrowRecord {
	var borrowRecord *entity.BorrowRecord
	if result, err := g.Model("borrow_record").Where(g.Map{"id": id}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&borrowRecord); err != nil {
		panic(err)
	}
	return borrowRecord
}

func (*BorrowRecordDaoImpl) SelectRecordsRestricted(record entity.BorrowRecord) []entity.BorrowRecord {
	records := []entity.BorrowRecord{}
	if result, err := g.Model("borrow_record").OmitEmpty().Where(map[string]any{
		"user_id":   record.UserId,
		"book_uuid": record.BookUUID,
		"state":     record.State,
	}).All(); err != nil {
		panic(err)
	} else {
		if err := result.Structs(&records); err != nil {
			panic(err)
		}
	}
	return records
}

func (*BorrowRecordDaoImpl) SelectRecordsInDetail(record entity.BorrowRecord, page do.QueryPage) do.PageData[do.BorrowRecordDetail] {
	var records = make([]do.BorrowRecordDetail, 0)
	var total int
	var err = g.Model(records, "user").OmitEmpty().Where(map[string]any{
		"user_id":   record.UserId,
		"book_uuid": record.BookUUID,
		"state":     record.State,
	}).Order("id desc").Limit((page.PageNum-1)*(page.PageSize), page.PageSize).WithAll().ScanAndCount(&records, &total, false)
	if err != nil {
		panic(err)
	}
	for _, v := range records {
		g.Dump(v.User)
	}
	return do.PageData[do.BorrowRecordDetail]{
		List:     records,
		PageNum:  page.PageNum,
		PageSize: page.PageSize,
		Total:    total,
	}
}
