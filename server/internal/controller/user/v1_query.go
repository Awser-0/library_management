package user

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) UserQuery() func(r *ghttp.Request) {
	type Form struct {
		QueryString string `json:"query_string" v:""`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var users = c.userService.UserQuery(form.QueryString)
		result.OK.SetMsg("登录成功").SetData(g.Map{
			"users": usersHandle(users),
		}).ToWriteJson(r)
	}
}
