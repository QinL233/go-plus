package app

import (
	"github.com/QinL233/go-plus/cache/redis"
	"github.com/QinL233/go-plus/mq/rocket"
	"github.com/QinL233/go-plus/orm/elastic"
	"github.com/QinL233/go-plus/orm/mysql"
	"github.com/QinL233/go-plus/oss/minio"
	"github.com/QinL233/go-plus/web"
	"github.com/QinL233/go-plus/yaml"
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
