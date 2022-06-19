package service

import "github.com/gogf/gf/v2/net/ghttp"

type IMiddleware interface {
	MiddlewareCORS(r *ghttp.Request)
	MiddlewareCORSDefault(r *ghttp.Request)
}

type middlewareImpl struct{}

var middleService = middlewareImpl{}

func Middleware() IMiddleware {
	return &middleService
}

func (s *middlewareImpl) MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	corsOptions.AllowDomain = []string{"*"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func (s *middlewareImpl) MiddlewareCORSDefault(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
