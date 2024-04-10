package user

import (
	"library/internal/model/do"
	"library/internal/utils"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Login() func(r *ghttp.Request) {
	type Form struct {
		Username string `json:"username" v:"required|length:2,10#请输入username|username长度为2-10"`
		Password string `json:"password" v:"required#请输入password"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var user, res = c.userService.Login(form.Username, form.Password)
		if res != nil {
			res.ToWriteJson(r)
			return
		}
		token, err := utils.GenerateTokenUsingHs256(do.AuthUser{
			UserID:  user.Id,
			IsAdmin: user.IsAdmin,
		})

		if err != nil {
			panic(gerror.Wrap(err, "token 生成失败"))
		}
		result.OK.SetMsg("登录成功").SetData(g.Map{
			"user":  user,
			"token": token,
		}).ToWriteJson(r)
	}
}
