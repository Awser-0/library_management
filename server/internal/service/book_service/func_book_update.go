package book_service

import (
	"library/internal/model/entity"
)

func (s *BookService) BookUpdate(book entity.Book) bool {
	return s.bookDao.UpdateBook(book.UUID, book)
}
