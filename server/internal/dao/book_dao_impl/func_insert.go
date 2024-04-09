package bookdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BookDaoImpl) InsertBook(book entity.Book) bool {
	result, err := g.Model("book").FieldsEx("uuid", "update_time", "create_time").Data(book).Insert()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
