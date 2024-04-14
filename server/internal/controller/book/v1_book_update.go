package book

import (
	"library/internal/utils"
	"library/internal/utils/result"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) BookUpdate() func(r *ghttp.Request) {
	type Form struct {
		UUID         int64      ` json:"uuid" v:"required"`
		Title        string     ` json:"title" v:"required"`
		Author       string     ` json:"author" v:"required"`
		Cover        string     ` json:"cover"`
		Introduction string     ` json:"introduction"`
		PublishTime  *time.Time ` json:"publish_time"`
		Total        int        ` json:"total"`
	}
	return func(r *ghttp.Request) {
		var form = utils.RequestParseForm[Form](r)
		var book = c.bookService.BookSelect(form.UUID)
		if book == nil {
			result.BookNotFound.ToWriteJson(r)
			return
		}
		_, form.Cover = filepath.Split(form.Cover)
		book.Title = form.Title
		book.Author = form.Author
		book.Cover = form.Cover
		book.Introduction = form.Introduction
		book.PublishTime = form.PublishTime
		book.Total = form.Total
		if c.bookService.BookUpdate(*book) {
			result.OK.SetMsg("修改成功").ToWriteJson(r)
		} else {
			result.Fail.SetMsg("修改失败").ToWriteJson(r)
		}
	}
}
