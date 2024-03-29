package web

import (
	"github.com/gin-gonic/gin"
	"reflect"
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
		c.Abort()
		return
	}
	c.JSON(200, parserResponse(data))
	c.Abort()
	return
}

func Fail(c *gin.Context, code int, err error) {
	if hadWriter(c) {
		c.Abort()
		return
	}
	c.JSON(200, Response{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	c.Abort()
	return
}

// 如果response已有东西则不再重复写
func hadWriter(c *gin.Context) bool {
	return c.Writer.Size() > 0
}

// 判断是否需要包装json
func parserResponse(data any) any {
	t := reflect.TypeOf(data)
	if t == nil {
		return Response{200, "success", nil}
	}
	if t.Kind() != reflect.Struct {
		return Response{200, "success", data}
	}
	_, d1 := t.FieldByName("Code")
	if !d1 {
		return Response{200, "success", data}
	}
	_, d2 := t.FieldByName("Msg")
	if !d2 {
		return Response{200, "success", data}
	}
	_, d3 := t.FieldByName("Data")
	if !d3 {
		return Response{200, "success", data}
	}
	return data
}
