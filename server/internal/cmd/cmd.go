package cmd

import (
	"context"
	"library/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				s.Group("/v1/user", func(group *ghttp.RouterGroup) {
					var c = user.NewV1()
					group.POST("/login", c.Login)
					group.POST("/register", c.Register)
				})
			})
			s.Run()
			return nil
		},
	}
)
