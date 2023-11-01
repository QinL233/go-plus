package web

import (
	"github.com/QinL233/go-plus/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
				if log.Driver() != nil {
					log.Driver().WithFields(logrus.Fields{
						"client": c.ClientIP(),
						"code":   c.Writer.Status(),
						"header": c.Request.Header,
						"method": c.Request.Method,
						"url":    c.Request.URL.Path,
					}).Errorf("panic: %v", err)
				}
				Fail(c, 500, err.(error))
			}
		}()
		c.Next()
		if len(c.Errors) > 0 {
			Fail(c, 500, c.Errors[0].Err)
		}
	}
}
