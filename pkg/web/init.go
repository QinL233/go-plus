package web

import (
	"fmt"
	"github.com/QinL233/go-plus/pkg/yaml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	if yaml.Config.Web.Port == 0 {
		yaml.Config.Web.Port = 8080
	}
	if yaml.Config.Web.Mode == "" {
		yaml.Config.Web.Mode = "info"
	}
	if yaml.Config.Web.ReadTimeout.String() == "0s" {
		yaml.Config.Web.ReadTimeout = time.Second * 60
	}
	if yaml.Config.Web.WriteTimeout.String() == "0s" {
		yaml.Config.Web.WriteTimeout = time.Second * 60
	}
}

// Init 对外提供初始化路由的方法
func Init(f func(r *gin.Engine)) {
	//级别
	gin.SetMode(yaml.Config.Web.Mode)
	r := gin.New()
	//默认日志
	r.Use(gin.Logger())
	//gin默认异常处理
	//r.Use(gin.Recovery())
	//自定义全局异常处理
	r.Use(recovery())

	//自定义拦截器
	for _, interceptor := range interceptors {
		r.Use(interceptor)
	}

	//自定义路由表
	router := r.Group(yaml.Config.Web.Prefix)
	for _, controller := range controllers {
		controller(router)
	}

	//自定义方法
	if f != nil {
		f(r)
	}

	//创建服务并启动
	log.Printf("start gin http server listening port[%d]", yaml.Config.Web.Port)
	err := (&http.Server{
		Addr:           fmt.Sprintf(":%d", yaml.Config.Web.Port),
		Handler:        r,
		ReadTimeout:    yaml.Config.Web.ReadTimeout,
		WriteTimeout:   yaml.Config.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}).ListenAndServe()
	if err != nil {
		log.Fatalf("gin server err %v", err)
	}
}
