package book

import (
	"library/internal/model/do"
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) RecordQuery() func(r *ghttp.Request) {
	type Form struct {
		BookUUID *int64 `json:"book_uuid" v:""`
		UserID   *int64 `json:"user_id" v:""`
		do.QueryPage
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var data do.PageData[do.BorrowRecordDetail]
		if form.BookUUID != nil {
			data = c.bookService.RecordQueryByBookUUID(*form.BookUUID, form.QueryPage)
		} else if form.UserID != nil {
			data = c.bookService.RecordQueryByUserID(*form.UserID, form.QueryPage)
		} else {
			data = c.bookService.RecordQuery(form.QueryPage)
		}
		result.OK.SetData(data).ToWriteJson(r)
	}
}
