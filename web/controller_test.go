package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"testing"
)

//1、构建 class 请求 param 和响应 result
type DemoParam struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file"`
}

type DemoResult struct {
	Password string `json:"password"`
	Ids      []int  `json:"ids"`
}

//2、定义service接口并实现
type DemoService struct {
	BaseService
	DemoParam
}

//通过实现exec使得controller能够回调
func (s *DemoService) Exec() (any, error) {
	return s.get(&s.DemoParam)
}

//真正的方法
func (s *DemoService) get(param *DemoParam) (*DemoResult, error) {
	fmt.Printf("%v", param)
	r := DemoResult{
		Password: "123456",
	}
	return &r, nil
}

//3、定义controller
func DemoController(c *gin.Context) {
	//使用controller工具快速构建响应体
	(&Controller[*DemoService]{Context: c}).Form()
}

func TestController(t *testing.T) {
	Router(func(g *gin.RouterGroup) {
		g.POST("/demo", DemoController)
	})
}
