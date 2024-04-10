package book_service

import "library/internal/dao"

type BookService struct {
	bookDao         dao.IBookDao
	borrowRecordDao dao.IBorrowRecordDao
}

func New() *BookService {
	return &BookService{
		bookDao:         dao.NewBookDao(),
		borrowRecordDao: dao.NewBorrowRecordDao(),
	}
}
