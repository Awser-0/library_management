package bookdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BookDaoImpl) UpdateBook(uuid int64, book entity.Book) bool {
	n, err := g.Model("book").FieldsEx("uuid", "update_time", "create_time").Data(book).Where("uuid", uuid).UpdateAndGetAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
