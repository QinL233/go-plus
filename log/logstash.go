package log

import (
	"bytes"
	"github.com/QinL233/go-plus/web"
	"github.com/QinL233/go-plus/yaml"
	logstash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net"
)

/**
日志记录到logstash中
*/

var driver *logrus.Logger

func initLogstash() {
	config := yaml.Config.Logstash
	if config.Url == "" {
		log.Println("logstash config is empty")
		return
	}
	driver = logrus.New()
	logMode := logrus.InfoLevel
	if yaml.Config.Web.Mode != gin.ReleaseMode {
		logMode = logrus.DebugLevel
	}
	driver.SetLevel(logMode)
	conn, err := net.Dial("tcp", config.Url)
	if err != nil {
		driver.Fatal(err)
	}
	hook := logstash.New(conn, logstash.DefaultFormatter(logrus.Fields{"tag": config.Tag}))

	driver.Hooks.Add(hook)

	log.Println("logstash connect success!")
	//添加一个web拦截器用于记录日志
	web.Interceptor(logstashInterceptor)
}

func logstashInterceptor(c *gin.Context) {
	logClient := driver.WithFields(logrus.Fields{
		"client": c.ClientIP(),
		"code":   c.Writer.Status(),
		"header": c.Request.Header,
		"method": c.Request.Method,
		"url":    c.Request.URL.Path,
	})
	if c.ContentType() == binding.MIMEJSON {
		body, err := c.GetRawData()
		if err != nil {
			go logClient.Printf("logstash read web request body err %v", err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		go logClient.Info(string(body))
	} else {
		go logClient.Info("")
	}
	c.Next()
}

func Driver() *logrus.Logger {
	return driver
}
