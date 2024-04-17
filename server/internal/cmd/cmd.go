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

				api.Group("/v1/user", func(r *ghttp.RouterGroup) {
					var c = user.NewV1()
					var router = r.Clone().Middleware()
					var userRouter = r.Clone().Middleware(middleware.MiddlewareAuth(false))
					var adminRouter = r.Clone().Middleware(middleware.MiddlewareAuth(true))
					{
						router.POST("/login", c.Login())
					}
					{
						userRouter.POST("/login/token", c.LoginByToken())
						userRouter.POST("/update", c.Update())
					}
					{
						adminRouter.POST("/register", c.Register())
						adminRouter.POST("/update/other", c.UpdateOther())
						adminRouter.POST("/query", c.UserQuery())
						adminRouter.POST("/select", c.UserSelect())
						adminRouter.POST("/pass/reset", c.PassReset())
					}
				})

				api.Group("/v1/book", func(r *ghttp.RouterGroup) {
					var c = book.NewV1()
					var router = r.Clone().Middleware()
					var userRouter = r.Clone().Middleware(middleware.MiddlewareAuth(false))
					var adminRouter = r.Clone().Middleware(middleware.MiddlewareAuth(true))
					{
						router.POST("/query", c.BookQuery())
						router.POST("/select", c.BookSelect())
					}
					{
						userRouter.POST("/borrow/apply", c.BorrowApply())
						userRouter.POST("/borrow/cancel", c.BorrowCancel())
						userRouter.POST("/record/query/self", c.RecordQuerySelf())
					}
					{
						adminRouter.POST("/add", c.BookAdd())
						adminRouter.POST("/update", c.BookUpdate())
						adminRouter.POST("/borrow/agree", c.BorrowAgree())
						adminRouter.POST("/borrow/reject", c.BorrowReject())
						adminRouter.POST("/borrow/return", c.BorrowReturn())
						adminRouter.POST("/record/query", c.RecordQuery())
					}
				})
			})
			s.Run()
			return nil
		},
	}
)
