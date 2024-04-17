package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) RecordQuerySelf() func(r *ghttp.Request) {
	type Form struct {
		BookUUID *int64 `json:"book_uuid" v:""`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var records = c.bookService.RecordQuerySelfByBookUUID(authUser.UserID, form.BookUUID)
		result.OK.SetData(g.Map{"records": records}).ToWriteJson(r)
	}
}
