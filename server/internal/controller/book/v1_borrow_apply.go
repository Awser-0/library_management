package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BorrowApply() func(r *ghttp.Request) {
	type Form struct {
		BookUUID int64  `json:"book_uuid" v:"required"`
		Desc     string `json:"desc" v:""`
	}
	return func(r *ghttp.Request) {
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var form = utils.RequestParseForm[Form](r)
		var ok, res = c.bookService.BorrowApply(form.BookUUID, authUser.UserID, form.Desc)
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
