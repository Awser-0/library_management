package cmd

import (
	"context"
	"library/internal/controller/book"
	"library/internal/controller/user"
	"library/internal/middleware"
	"library/internal/utils/result"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func in(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			var ImageSavePath = "C:\\work\\images"
			s.Use(middleware.MiddlewareCORS)
			s.AddStaticPath("/images", ImageSavePath)
			s.Use(middleware.MiddlewareException)
			s.Group("/api", func(api *ghttp.RouterGroup) {

				api.Group("/v1/file", func(group *ghttp.RouterGroup) {
					var adminRouter = group.Middleware(middleware.MiddlewareAuth(true))
					adminRouter.POST("/upload/image", func(r *ghttp.Request) {
						file := r.GetUploadFile("upload-file")
						if file == nil {
							result.Fail.SetMsg("上传失败").ToWriteJson(r)
							return
						}
						if file.Size/1024/1024 >= 20 {
							result.Fail.SetMsg("图片大小不能超过20MB").ToWriteJson(r)
							return
						}
						if !in(strings.ToLower(path.Ext(file.Filename)), []string{".png", ".jpg", ".jpeg"}) {
							result.Fail.SetMsg("图片类型仅支持png、jpg、jpeg").ToWriteJson(r)
							return
						}
						file.Filename = strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)
						name, err := file.Save(ImageSavePath)
						if err != nil {
							result.Fail.SetMsg(err.Error()).ToWriteJson(r)
						} else {
							result.OK.SetData(g.Map{"name": name}).ToWriteJson(r)
						}
					})
				})

				api.Group("/v1/user", func(v1user *ghttp.RouterGroup) {
					var c = user.NewV1()
					v1user.POST("/login", c.Login())
					var userRouter = v1user.Middleware(middleware.MiddlewareAuth(false))
					userRouter.POST("/login/token", c.LoginByToken())
					var adminRouter = v1user.Middleware(middleware.MiddlewareAuth(true))
					adminRouter.POST("/register", c.Register())
					adminRouter.POST("/query", c.UserQuery())
				})

				api.Group("/v1/book", func(v1book *ghttp.RouterGroup) {
					var c = book.NewV1()
					{
						v1book.POST("/query", c.BookQuery())
						v1book.POST("/select", c.BookSelect())
					}
					var userRouter = v1book.Middleware(middleware.MiddlewareAuth(false))
					{
						userRouter.POST("/borrow/apply", c.BorrowApply())
					}
					var adminRouter = v1book.Middleware(middleware.MiddlewareAuth(true))
					{
						adminRouter.POST("/add", c.BookAdd())
						adminRouter.POST("/update", c.BookUpdate())
						adminRouter.POST("/borrow/agree", c.BorrowAgree())
						adminRouter.POST("/borrow/reject", c.BorrowReject())
						adminRouter.POST("/borrow/return", c.BorrowReturn())
						adminRouter.POST("/record/list", c.RecordList())
					}
				})
			})
			s.Run()
			return nil
		},
	}
)
