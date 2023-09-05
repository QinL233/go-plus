package app

import (
	"go-plus/cache/redis"
	"go-plus/gin"
	"go-plus/mq/rocket"
	"go-plus/orm/elastic"
	"go-plus/orm/mysql"
	"go-plus/oss/minio"
	"go-plus/yaml"
)

func Start(conf ...string) {
	yaml.Init(conf...)
	mysql.Init()
	elastic.Init()
	redis.Init()
	minio.Init()
	rocket.Init()
	gin.Init()
}
