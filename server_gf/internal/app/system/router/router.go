package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"server_gf/internal/app/system/controller"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		//绑定路由:系统参数
		group.Bind(controller.SysConfig)
	})
}
