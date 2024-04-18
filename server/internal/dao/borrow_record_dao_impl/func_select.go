package borrowrecorddaoimpl

import (
	"fmt"
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
	}).Where("").All(); err != nil {
		panic(err)
	} else {
		if err := result.Structs(&records); err != nil {
			panic(err)
		}
	}
	return records
}

func (*BorrowRecordDaoImpl) SelectRecordsInDetail() []do.BorrowRecordDetail {
	var records []do.BorrowRecordDetail
	fmt.Printf("records: %v\n", records)
	if err := g.Model(records).WithAll().Scan(&records); err != nil {
		panic(err)
	}
	return records
}
