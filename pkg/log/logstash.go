package log

import (
	"github.com/QinL233/go-plus/pkg/yaml"
	logstash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	logMode := logrus.InfoLevel
	if yaml.Config.Web.Mode != gin.ReleaseMode {
		logMode = logrus.DebugLevel
	}
	driver = logrus.New()
	driver.SetLevel(logMode)
	conn, err := net.Dial("tcp", config.Url)
	if err != nil {
		log.Fatalf("connect logstash err: %v", err)
	}
	hook := logstash.New(conn, logstash.DefaultFormatter(logrus.Fields{"tag": config.Tag}))

	driver.Hooks.Add(hook)

	log.Println("logstash connect success!")
}

func Driver() *logrus.Logger {
	return driver
}
