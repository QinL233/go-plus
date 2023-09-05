package app

import (
	"go-plus/cache/redis"
	"go-plus/gin"
	"go-plus/orm/mysql"
	"go-plus/oss/minio"
	"go-plus/yaml"
)

func Start(conf ...string) {
	yaml.Init(conf...)
	mysql.Init()
	redis.Init()
	minio.Init()
	gin.Init()
}
