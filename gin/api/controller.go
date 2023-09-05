package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"log"
)

/**
参数解析器
1、分组各类传参
2、传参校验
3、根据传参转发到指定handler
4、封装返回
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
