package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/protocol/goai"
	"server_gf/internal/consts"
	"server_gf/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//使用中间件
			s.Use(ghttp.MiddlewareHandlerResponse)
			//绑定路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				//跨域
				group.Middleware(
					ghttp.MiddlewareCORS,
				)
				//绑定处理器
				router.BindController(group)
			})
			enhanceOpenAPIDoc(s)
			//s.SetFileServerEnabled(true)
			//s.SetIndexFolder(true)
			//s.AddSearchPath("D:\\study\\dev\\gf2-vue-admin\\web\\dist")
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
}
