package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BookSelect() func(r *ghttp.Request) {
	type Form struct {
		UUID int64 ` json:"uuid" v:"required"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var book = c.bookService.BookSelect(form.UUID)
		if book == nil {
			result.BookNotFound.ToWriteJson(r)
			return
		} else {
			result.OK.SetData(g.Map{"book": book}).ToWriteJson(r)
			return
		}
	}
}
