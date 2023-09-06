package web

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
Service - 工具
1、用于快速从context实现做提取param并校验
2、封装service.func(param)(result,err)方法的返回
*/

func Service[P any, R any](c *gin.Context, f func(P) (R, error)) {
	var param P
	if c.Params != nil && len(c.Params) > 0 {
		//uri请求 - 必须注意service使用`uri:"paramName"`接收
		if err := c.ShouldBindUri(&param); err != nil {
			log.Println(err)
			fail(c, 500, errors.New("param is fail"))
			return
		}
	} else {
		//query/json/form请求
		if err := c.ShouldBind(&param); err != nil {
			log.Println(err)
			fail(c, 500, errors.New("param is fail"))
			return
		}
	}
	r, err := f(param)
	if err != nil {
		fail(c, 500, err)
		return
	}
	success(c, r)
}

/**
封装标准的返回格式
*/

type JSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func success(c *gin.Context, data any) {
	c.JSON(200, JSON{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	return
}

func fail(c *gin.Context, code int, err error) {
	c.JSON(200, JSON{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	return
}
