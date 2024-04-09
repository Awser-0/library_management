package dao

import (
	bookdaoimpl "library/internal/dao/book_dao_impl"
	"library/internal/model/entity"
)

type IBookDao interface {
	SelectBook(uuid int64) *entity.Book
	SelectBooksRestricted(queryString string) []entity.Book
	InsertBook(book entity.Book) bool
	UpdateBook(uuid int64, book entity.Book) bool
	DeleteBook(uuid int64) bool
}

func NewBookDao() IBookDao {
	return bookdaoimpl.New()
}
