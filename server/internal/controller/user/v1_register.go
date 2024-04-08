package user

import (
	httpException "library/internal/utils/http_exception"
	"library/internal/utils/result"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type v1RegisterDTO struct {
	Username string `json:"username" v:"required|length:2,10#请输入username|username长度为2-10"`
	Password string `json:"password" v:"required#请输入password"`
}

func (c *ControllerV1) Register(r *ghttp.Request) {
	var form *v1RegisterDTO
	if err := r.Parse(&form); err != nil {
		r.SetError(httpException.BadRequestHttpException.SetMsg(err.Error()))
		return
	}
	r.Response.WriteJson(result.OK.SetMsg("登录成功").SetData(g.Map{
		"form": form,
	}).ToMap())
}
