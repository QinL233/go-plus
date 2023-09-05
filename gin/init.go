package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-plus/gin/api"
	"go-plus/gin/interceptor"
	"go-plus/yaml"
	"log"
	"net/http"
	"time"
)

func init() {
	if yaml.Config.Gin.Port == 0 {
		yaml.Config.Gin.Port = 8080
	}
	if yaml.Config.Gin.Mode == "" {
		yaml.Config.Gin.Mode = "info"
	}
	if yaml.Config.Gin.ReadTimeout.String() == "0s" {
		yaml.Config.Gin.ReadTimeout = time.Second * 60
	}
	if yaml.Config.Gin.WriteTimeout.String() == "0s" {
		yaml.Config.Gin.WriteTimeout = time.Second * 60
	}
}

// Init 对外提供初始化路由的方法
func Init() {
	r := gin.New()
	//默认日志
	r.Use(gin.Logger())
	//500
	r.Use(gin.Recovery())

	//自定义拦截器
	for _, interceptor := range interceptor.Tables {
		r.Use(interceptor)
	}

	//自定义路由表
	router := r.Group(yaml.Config.Gin.Prefix)
	for _, controller := range api.Tables {
		controller(router)
	}

	//级别
	gin.SetMode(yaml.Config.Gin.Mode)

	//创建服务并启动
	log.Printf("start gin http server listening port[%d]", yaml.Config.Gin.Port)
	err := (&http.Server{
		Addr:           fmt.Sprintf(":%d", yaml.Config.Gin.Port),
		Handler:        r,
		ReadTimeout:    yaml.Config.Gin.ReadTimeout,
		WriteTimeout:   yaml.Config.Gin.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}).ListenAndServe()
	if err != nil {
		log.Fatalf("gin server err %v", err)
	}
}
