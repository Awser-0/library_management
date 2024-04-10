package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BorrowReturn() func(r *ghttp.Request) {
	type Form struct {
		RecordId int64 `json:"record_id" v:"required"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var ok, res = c.bookService.BorrowReturn(form.RecordId)
		if res != nil {
			res.ToWriteJson(r)
		} else {
			if ok {
				result.OK.ToWriteJson(r)
			} else {
				result.Fail.ToWriteJson(r)
			}
		}
	}
}
