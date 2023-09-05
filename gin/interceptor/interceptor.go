package interceptor

import "github.com/gin-gonic/gin"

/**
拦截器表
*/
var Tables = make([]func(g *gin.Context), 0)

// Interceptor 用于注册拦截器
func Interceptor(f func(g *gin.Context)) {
	Tables = append(Tables, f)
}
