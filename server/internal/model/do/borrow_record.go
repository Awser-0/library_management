package do

import (
	"library/internal/model/entity"
)

type BorrowRecordDetail struct {
	entity.BorrowRecord
	User *entity.User `json:"user" orm:"with:id=user_id"`
	Book *entity.Book `json:"book" orm:"with:uuid=book_uuid"`
}
