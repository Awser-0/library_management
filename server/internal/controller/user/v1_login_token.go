package user

import (
	"library/internal/utils"
	httpException "library/internal/utils/http_exception"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) LoginByToken() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		var authUser = utils.GetAuthUserFromCtx(r.GetCtx())
		var user = c.userService.SelectUser(authUser.UserID)
		if user == nil {
			r.SetError(httpException.UnauthorizedHttpException)
			return
		}
		result.OK.SetMsg("登录成功").SetData(g.Map{
			"user": user,
		}).ToWriteJson(r)
	}
}
