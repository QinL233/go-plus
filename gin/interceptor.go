package gin

import "github.com/gin-gonic/gin"

/**
拦截器表
*/
var interceptorTables = make([]func(g *gin.Context), 0)

// Interceptor 用于注册拦截器
func Interceptor(f func(g *gin.Context)) {
	interceptorTables = append(interceptorTables, f)
}
