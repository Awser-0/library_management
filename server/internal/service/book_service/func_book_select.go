package book_service

import "library/internal/model/entity"

func (s *BookService) BookSelect(uuid int64) *entity.Book {
	return s.bookDao.SelectBook(uuid)
}
