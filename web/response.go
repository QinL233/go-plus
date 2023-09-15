package web

import (
	"github.com/gin-gonic/gin"
)

/**
封装标准的返回格式
*/

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	if hadWriter(c) {
		return
	}
	c.JSON(200, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	return
}

func Fail(c *gin.Context, code int, err error) {
	if hadWriter(c) {
		return
	}
	c.JSON(200, Response{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	return
}

func Json(c *gin.Context, data any) {
	if hadWriter(c) {
		return
	}
	c.JSON(200, data)
	return
}

//如果response已有东西则不再重复写
func hadWriter(c *gin.Context) bool {
	return c.Writer.Size() > 0
}
