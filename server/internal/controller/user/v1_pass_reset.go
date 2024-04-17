package user

import (
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) PassReset() func(r *ghttp.Request) {
	type Form struct {
		UserID   int64  `json:"user_id" v:"required"`
		Password string `json:"password" v:"required#请输入password"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var user = c.userService.SelectUser(form.UserID)
		if user == nil {
			result.UserNotFound.ToWriteJson(r)
			return
		}
		user.Password = form.Password
		var ok, res = c.userService.UpdateUser(*user)
		if res != nil {
			res.ToWriteJson(r)
			return
		}
		if ok {
			result.OK.SetMsg("修改成功").SetData(ok).ToWriteJson(r)
			return
		}
	}
}
