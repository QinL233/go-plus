package web

import (
	"github.com/gin-gonic/gin"
)

/**
路由表 - 用于注册controller的入口
*/

var controllers []func(g *gin.RouterGroup)

// Controller 用于controller注册路由
func Controller(list ...func(g *gin.RouterGroup)) {
	if controllers == nil {
		controllers = make([]func(g *gin.RouterGroup), 0)
	}
	for _, controller := range list {
		controllers = append(controllers, controller)
	}
}
