package book_service

import (
	"library/internal/model/entity"
	"time"
)

func (s *BookService) BookAdd(title, author, cover, introduction string, publishTime *time.Time, total int) bool {
	return s.bookDao.InsertBook(entity.Book{
		Title:        title,
		Author:       author,
		Cover:        cover,
		Introduction: introduction,
		PublishTime:  publishTime,
		Total:        total,
	})
}
