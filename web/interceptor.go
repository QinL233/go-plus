package web

import "github.com/gin-gonic/gin"

/**
拦截器表
*/
var interceptors []func(g *gin.Context)

// Interceptor 用于注册拦截器
func Interceptor(handlers ...func(g *gin.Context)) {
	if interceptors == nil {
		interceptors = make([]func(g *gin.Context), 0)
	}
	for _, handler := range handlers {
		interceptors = append(interceptors, handler)
	}
}
