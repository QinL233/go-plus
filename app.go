package app

import (
	"go-plus/cache/redis"
	"go-plus/mq/rocket"
	"go-plus/orm/elastic"
	"go-plus/orm/mysql"
	"go-plus/oss/minio"
	"go-plus/web"
	"go-plus/yaml"
)

/**
服务启动入口
*/

func Start(conf ...string) {
	yaml.Init(conf...)
	mysql.Init()
	elastic.Init()
	redis.Init()
	minio.Init()
	rocket.Init()
	web.Init()
}
