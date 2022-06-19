package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	systemRouter "server_gf/internal/app/system/router"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//绑定后台路由
		systemRouter.BindController(group)
	})
}
