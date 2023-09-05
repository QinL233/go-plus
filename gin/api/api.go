package api

import (
	"github.com/gin-gonic/gin"
)

/**
路由表
*/

var Tables = make([]func(g *gin.RouterGroup), 0)

// Api 用于controller注册路由
func Api(f func(g *gin.RouterGroup)) {
	Tables = append(Tables, f)
}
