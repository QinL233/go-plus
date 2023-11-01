package web

import (
	"fmt"
	"github.com/QinL233/go-plus/log"
	"github.com/QinL233/go-plus/orm/mysql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
快速构建controller的方法
1、快速定义api、type、入参、响应、处理函数
2、封装gin.context不进行传递使用
3、封装controller提取和校验入参(param)
4、封装driver的生成
5、规范化标准的server：传递driver便于事务、panic err以统一异常捕获
*/

func GET[P any, R any](uri string, server func(db *gorm.DB, param P) R) {
	Controller(func(g *gin.RouterGroup) {
		g.GET(uri, func(context *gin.Context) {
			Server[P, R](context, server)
		})
	})
}

func POST[P any, R any](uri string, server func(db *gorm.DB, param P) R) {
	Controller(func(g *gin.RouterGroup) {
		g.POST(uri, func(context *gin.Context) {
			Server[P, R](context, server)
		})
	})
}

func PUT[P any, R any](uri string, server func(db *gorm.DB, param P) R) {
	Controller(func(g *gin.RouterGroup) {
		g.PUT(uri, func(context *gin.Context) {
			Server[P, R](context, server)
		})
	})
}

func DELETE[P any, R any](uri string, server func(db *gorm.DB, param P) R) {
	Controller(func(g *gin.RouterGroup) {
		g.DELETE(uri, func(context *gin.Context) {
			Server[P, R](context, server)
		})
	})
}

// Server controller前后文处理（隐藏context，生成driver、规范化server函数）
func Server[P any, R any](c *gin.Context, f func(db *gorm.DB, param P) R) {
	var param P
	//1、【尝试】从header中取得参数【不保证校验】 - 获取`header:"paramName"`
	c.ShouldBindHeader(&param)
	//2、根据参数性质bind参数
	if c.Params != nil && len(c.Params) > 0 {
		//uri请求 - 必须注意param struct使用`uri:"paramName"`标签接收解析
		if err := c.ShouldBindUri(&param); err != nil {
			fmt.Println(err)
			Fail(c, 500, errors.New("param is fail"))
			return
		}
	} else {
		//query/json/form请求 - 必须注意param struct使用`form:"paramName"`标签接收解析
		if err := c.ShouldBind(&param); err != nil {
			fmt.Println(err)
			Fail(c, 500, errors.New("param is fail"))
			return
		}
	}
	//3.回调server方法并返回封装
	if log.Driver() != nil {
		//将输入输出打印到日志中
		r := f(mysql.Driver(), param)
		log.Driver().WithFields(logrus.Fields{
			"client": c.ClientIP(),
			"code":   c.Writer.Status(),
			"header": c.Request.Header,
			"method": c.Request.Method,
			"url":    c.Request.URL.Path,
			"param":  param,
			"result": r,
		}).Info()
		Success(c, r)
	} else {
		Success(c, f(mysql.Driver(), param))
	}
}
