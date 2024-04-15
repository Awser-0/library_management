package user

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) UserSelect() func(r *ghttp.Request) {
	type Form struct {
		Id int64 `json:"id" v:"required"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var user = c.userService.SelectUser(form.Id)
		if user == nil {
			result.UserNotFound.ToWriteJson(r)
			return
		} else {
			result.OK.SetData(g.Map{"user": userHandle(*user)}).ToWriteJson(r)
			return
		}
	}
}
