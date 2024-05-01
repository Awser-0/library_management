package user

import (
	"library/internal/model/do"
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) UserQuery() func(r *ghttp.Request) {
	type Form struct {
		QueryString string `json:"query_string" v:""`
		do.QueryPage
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var data = c.userService.UserQuery(form.QueryString, form.QueryPage)
		result.OK.SetData(data).ToWriteJson(r)
	}
}
