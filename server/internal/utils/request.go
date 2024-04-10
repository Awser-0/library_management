package utils

import (
	httpException "library/internal/utils/http_exception"

	"github.com/gogf/gf/v2/net/ghttp"
)

func RequestParseForm[T any](r *ghttp.Request) *T {
	var form *T
	if err := r.Parse(&form); err != nil {
		r.SetError(httpException.BadRequestHttpException.SetMsg(err.Error()))
		r.Exit()
	}
	return form
}
