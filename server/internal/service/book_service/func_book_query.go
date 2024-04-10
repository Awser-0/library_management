package book_service

import "library/internal/model/entity"

func (s *BookService) BookQuery(queryString string) []entity.Book {
	return s.bookDao.SelectBooksRestricted(queryString)
}
