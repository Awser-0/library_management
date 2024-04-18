package test

import (
	"fmt"
	"library/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func Test_BorrowRecord(t *testing.T) {
	// var d = dao.NewBorrowRecordDao()
	// var records = d.SelectRecordsInDetail()
	// g.Dump(records)
	var users []entity.User
	var all, count, _ = g.Model("user").Limit(-2, -1).AllAndCount(false)
	all.Structs(&users)
	fmt.Printf("count: %v\n", count)
	g.Dump(users)
}
