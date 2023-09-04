package example

import (
	"fmt"
	app "go-plus"
	"go-plus/gin"
	"testing"
)

type testService struct {
	db string
}

func (s *testService) handler() {
	fmt.Println(fmt.Sprintf("test %p", s))
}

func (s *testService) handler2() {
	fmt.Println(fmt.Sprintf("test %p", s))
}

func TestHttp(t *testing.T) {
	gin.Router(
		&gin.Controller{
			Method:  gin.GET,
			Api:     "/demo",
			Service: (&testService{}).handler,
		},
		&gin.Controller{
			Method:  gin.GET,
			Api:     "/demo2",
			Service: (&testService{}).handler2,
		},
	)
	app.Start("config/app.yml")
}
