package book

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BookQuery() func(r *ghttp.Request) {
	type Form struct {
		QueryString string `json:"query_string" v:""`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var books = c.bookService.BookQuery(form.QueryString)
		result.OK.SetData(g.Map{
			"books": books,
		}).ToWriteJson(r)
	}
}
