package book

import (
	"library/internal/utils"
	"library/internal/utils/result"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

//

func (c *ControllerV1) BookAdd() func(r *ghttp.Request) {
	type Form struct {
		Title        string     ` json:"title" v:"required"`
		Author       string     ` json:"author" v:"required"`
		Cover        string     ` json:"cover"`
		Introduction string     ` json:"introduction"`
		PublishTime  *time.Time ` json:"publish_time"`
		Total        int        ` json:"total"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		_, form.Cover = filepath.Split(form.Cover)
		var ok = c.bookService.BookAdd(form.Title, form.Author, form.Cover, form.Introduction, form.PublishTime, form.Total)
		if ok {
			result.OK.ToWriteJson(r)
		} else {
			result.Fail.ToWriteJson(r)
		}
	}
}
