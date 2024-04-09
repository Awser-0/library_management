package borrowrecorddaoimpl

import "github.com/gogf/gf/v2/frame/g"

func (*BorrowRecordDaoImpl) DeleteRecord(id int64) bool {
	result, err := g.Model("borrow_record").Where("id", id).Delete()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
