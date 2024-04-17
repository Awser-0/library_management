package book

import (
	"library/internal/model/entity"
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) RecordQuery() func(r *ghttp.Request) {
	type Form struct {
		BookUUID *int64 `json:"book_uuid" v:""`
		UserID   *int64 `json:"user_id" v:""`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var records []entity.BorrowRecord
		if form.BookUUID != nil {
			records = c.bookService.RecordQueryByBookUUID(*form.BookUUID)
		} else if form.UserID != nil {
			records = c.bookService.RecordQueryByUserID(*form.UserID)
		} else {
			records = c.bookService.RecordQuery()
		}
		result.OK.SetData(g.Map{"records": records}).ToWriteJson(r)
	}
}
