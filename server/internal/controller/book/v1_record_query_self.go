package book

import (
	"library/internal/model/do"
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) RecordQuerySelf() func(r *ghttp.Request) {
	type Form struct {
		BookUUID *int64 `json:"book_uuid" v:""`
		do.QueryPage
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var data = c.bookService.RecordQuerySelfByBookUUID(authUser.UserID, form.BookUUID, form.QueryPage)
		result.OK.SetData(data).ToWriteJson(r)
	}
}
