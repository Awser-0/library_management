package bookdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*BookDaoImpl) SelectBook(uuid int64) *entity.Book {
	var book *entity.Book
	if result, err := g.Model("book").Where(g.Map{"uuid": uuid}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&book); err != nil {
		panic(err)
	}
	return book
}

func (*BookDaoImpl) SelectBooksRestricted(queryString string) []entity.Book {
	queryString = queryString + "%"
	var books = []entity.Book{}
	if result, err := g.Model("book").WhereLike("title", queryString).WhereOrLike("author", queryString).All(); err != nil {
		panic(err)
	} else {
		if err := result.Structs(&books); err != nil {
			panic(err)
		}
	}
	return books
}
