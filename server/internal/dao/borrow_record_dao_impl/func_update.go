package borrowrecorddaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BorrowRecordDaoImpl) UpdateRecord(id int64, book entity.BorrowRecord) bool {
	n, err := g.Model("borrow_record").FieldsEx("id", "update_time", "create_time").Data(book).Where("id", id).UpdateAndGetAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
