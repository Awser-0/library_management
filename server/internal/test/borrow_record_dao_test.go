package test

import (
	"fmt"
	"library/internal/dao"
	"library/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_BorrowRecordDao_Select(t *testing.T) {
	var dao = dao.NewBorrowRecordDao()
	// gtest.C(t, func(t *gtest.T) {
	// 	var record = dao.SelectRecord(1)
	// 	fmt.Printf("record: %v\n", record)
	// 	t.Assert(record != nil, true)
	// })
	gtest.C(t, func(t *gtest.T) {
		var records = dao.SelectRecordsRestricted(entity.BorrowRecord{})
		for i, r := range records {
			fmt.Printf("%v: %v\n", i, r)
		}
		t.Assert(len(records) != 0, true)
	})
}

func Test_BorrowRecordDao_Insert(t *testing.T) {
	var dao = dao.NewBorrowRecordDao()
	gtest.C(t, func(t *gtest.T) {
		t.Assert(dao.InsertRecord(entity.BorrowRecord{
			UserId:     1,
			BookUUID:   10000002,
			State:      entity.BorrowRecordStateApply,
			ReplyTime:  nil,
			ReturnTime: nil,
			ApplyDesc:  "想看",
			ReplyDesc:  "",
		}), true)
	})
}

func Test_BorrowRecordDao_Update(t *testing.T) {
	var dao = dao.NewBorrowRecordDao()
	gtest.C(t, func(t *gtest.T) {
		var record = dao.SelectRecord(1)
		record.ApplyDesc = "真的想看2"
		t.Assert(dao.UpdateRecord(record.Id, *record), true)
	})
}

func Test_BorrowRecordDao_Delete(t *testing.T) {
	var dao = dao.NewBorrowRecordDao()
	gtest.C(t, func(t *gtest.T) {
		t.Assert(dao.DeleteRecord(11), true)
	})
}
