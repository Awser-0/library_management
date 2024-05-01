package user

import (
	"library/internal/utils"
	"library/internal/utils/result"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Update() func(r *ghttp.Request) {
	type Form struct {
		Username string     `json:"username" v:"required|length:2,10#请输入username|username长度为2-10"`
		Nickname string     `json:"nickname" v:"required"`
		Sex      string     `json:"sex" v:"required"`
		Phone    string     `json:"phone" v:""`
		Birth    *time.Time `json:"birth" v:""`
		IsAdmin  bool       `json:"isAdmin" v:"required"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var user = c.userService.SelectUser(authUser.UserID)
		if user == nil {
			result.UserNotFound.ToWriteJson(r)
			return
		}
		user.Username = form.Username
		user.Nickname = form.Nickname
		user.Sex = form.Sex
		user.Phone = form.Phone
		user.Birth = form.Birth
		user.IsAdmin = form.IsAdmin
		var ok, res = c.userService.UpdateUser(*user)
		if res != nil {
			res.ToWriteJson(r)
			return
		}
		if ok {
			result.OK.SetMsg("修改成功").ToWriteJson(r)
			return
		} else {
			result.Fail.SetMsg("修改失败").ToWriteJson(r)
			return
		}
	}
}
