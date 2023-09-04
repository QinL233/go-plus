package gin

import (
	"github.com/gin-gonic/gin"
)

/**
路由表
*/
//var routerTables = make([]func(g *gin.RouterGroup), 0)
//
//// Router 用于注册路由
//func Router(f func(g *gin.RouterGroup)) {
//	routerTables = append(routerTables, f)
//}

const (
	GET = iota
	POST
	PUT
	DELETE
)

type Controller struct {
	*gin.Context
	Method       int         //请求类型
	Api          string      //请求api
	Param        interface{} //请求参数
	ParamHandler func()      //参数处理函数
	Service      Service     //service方法
}

type Service func()

var routerTables = make([]*Controller, 0)

// Router 用于注册路由
func Router(cs ...*Controller) {
	for _, c := range cs {
		routerTables = append(routerTables, c)
	}
}

func InitRouter(g *gin.Engine, pre string) {
	r := g.Group(pre)
	for _, c := range routerTables {
		if c.Method == GET {
			r.GET(c.Api, func(context *gin.Context) {
				c.Context = context
				c.Service()
			})
		}
	}
}
