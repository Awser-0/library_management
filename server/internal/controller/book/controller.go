package book

import (
	"library/internal/service"
)

type ControllerV1 struct {
	bookService service.IBookService
}

func NewV1() ControllerV1 {
	return ControllerV1{
		bookService: service.NewBookService(),
	}
}
