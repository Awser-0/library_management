package borrowrecorddaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BorrowRecordDaoImpl) InsertRecord(record entity.BorrowRecord) bool {
	result, err := g.Model("borrow_record").FieldsEx("id", "update_time", "create_time").Data(record).Insert()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
