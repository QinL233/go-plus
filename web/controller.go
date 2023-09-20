package web

import (
	"github.com/QinL233/go-plus/orm/mysql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
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

/**
快速构建controller的方法，用于快速的定义入参、响应、server函数
*/

func GET[P any, R any](uri string, f func(db *gorm.DB, p P) (r R, err error)) {
	Controller(func(g *gin.RouterGroup) {
		g.GET(uri, func(context *gin.Context) {
			Server[P, R](context, f)
		})
	})
}

func POST[P any, R any](uri string, f func(db *gorm.DB, p P) (r R, err error)) {
	Controller(func(g *gin.RouterGroup) {
		g.POST(uri, func(context *gin.Context) {
			Server[P, R](context, f)
		})
	})
}

func PUT[P any, R any](uri string, f func(db *gorm.DB, p P) (r R, err error)) {
	Controller(func(g *gin.RouterGroup) {
		g.PUT(uri, func(context *gin.Context) {
			Server[P, R](context, f)
		})
	})
}

func DELETE[P any, R any](uri string, f func(db *gorm.DB, p P) (r R, err error)) {
	Controller(func(g *gin.RouterGroup) {
		g.DELETE(uri, func(context *gin.Context) {
			Server[P, R](context, f)
		})
	})
}

func Server[P any, R any](c *gin.Context, f func(db *gorm.DB, param P) (R, error)) {
	var param P
	//1、尝试从header中取得参数【不保证校验】 - 获取`header:"paramName"`
	err := c.ShouldBindHeader(&param)
	if err != nil {
		log.Println(err)
	}
	//2、根据参数性质bind参数
	if c.Params != nil && len(c.Params) > 0 {
		//uri请求 - 必须注意service使用`uri:"paramName"`接收
		if err = c.ShouldBindUri(&param); err != nil {
			Fail(c, 500, errors.New("param is fail"))
			return
		}
	} else {
		//query/json/form请求
		if err = c.ShouldBind(&param); err != nil {
			Fail(c, 500, errors.New("param is fail"))
			return
		}
	}
	if err != nil {
		log.Println(err)
		Fail(c, 500, err)
		return
	}
	//3.回调方法
	r, err := f(mysql.Driver(), param)
	if err != nil {
		Fail(c, 500, err)
		return
	}
	//4.构建返回
	Success(c, r)
}
