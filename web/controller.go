package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"log"
)

/**
路由表 - 用于注册controller的入口
*/

var routers []func(g *gin.RouterGroup)

// Router 用于controller注册路由
func Router(controllers ...func(g *gin.RouterGroup)) {
	if routers == nil {
		routers = make([]func(g *gin.RouterGroup), 0)
	}
	for _, controller := range controllers {
		routers = append(routers, controller)
	}
}

/**
controller执行器
*/

type Controller[S Service] struct {
	*gin.Context
}

func (c *Controller[S]) Query() {
	var service S
	if err := c.Context.ShouldBindQuery(&service); err != nil {
		log.Println(err)
		fail(c.Context, 500, errors.New("param is fail"))
		return
	}
	c.handler(service)
}

func (c *Controller[S]) Path() {
	var service S
	if err := c.Context.ShouldBindUri(&service); err != nil {
		log.Println(err)
		fail(c.Context, 500, errors.New("param is fail"))
		return
	}
	c.handler(service)
}

func (c *Controller[S]) Body() {
	var service S
	if err := c.Context.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		fail(c.Context, 500, errors.New("param is fail"))
		return
	}
	c.handler(service)
}

func (c *Controller[S]) Form() {
	var service S
	if err := c.Context.ShouldBindWith(&service, binding.FormMultipart); err != nil {
		log.Println(err)
		fail(c.Context, 500, errors.New("param is fail"))
		return
	}
	c.handler(service)
}

func (c *Controller[S]) handler(service S) {
	r, err := service.handler(service)
	if err != nil {
		fail(c.Context, 500, err)
		return
	}
	success(c.Context, r)
}
