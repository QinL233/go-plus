package web

import "github.com/gin-gonic/gin"

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
