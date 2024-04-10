package user

import (
	"library/internal/utils"
	httpException "library/internal/utils/http_exception"
	"library/internal/utils/result"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Register() func(r *ghttp.Request) {
	type Form struct {
		Username string     `json:"username" v:"required|length:2,10#请输入username|username长度为2-10"`
		Password string     `json:"password" v:"required#请输入password"`
		Sex      string     `json:"sex" v:"required"`
		Phone    string     `json:"phone" v:"required"`
		Birth    *time.Time `json:"birth" v:"required"`
		IsAdmin  bool       `json:"isAdmin" v:"required"`
	}
	return func(r *ghttp.Request) {
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		if !authUser.IsAdmin {
			r.SetError(httpException.ForbiddenHttpException)
			return
		}
		var form = utils.RequestParseForm[Form](r)
		var ok, res = c.userService.Register(form.Username, form.Password, form.Sex, form.Phone, form.Birth, form.IsAdmin)
		if res != nil {
			res.ToWriteJson(r)
		} else {
			result.OK.SetData(ok).ToWriteJson(r)
		}
	}
}
