package web

import (
	"github.com/gin-gonic/gin"
)

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

//自定义异常拦截器
func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Fail(c, 500, err.(error))
			}
		}()
		c.Next()
		if len(c.Errors) > 0 {
			Fail(c, 500, c.Errors[0].Err)
		}
	}
}
