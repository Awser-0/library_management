package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BorrowCancel() func(r *ghttp.Request) {
	type Form struct {
		RecordId int64 `json:"record_id" v:"required"`
	}
	return func(r *ghttp.Request) {
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var form = utils.RequestParseForm[Form](r)
		var ok, res = c.bookService.BorrowCancel(form.RecordId, authUser.UserID)
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
