package web

import (
	"github.com/QinL233/go-plus/pkg/cache/redis"
	"github.com/QinL233/go-plus/pkg/cron"
	"github.com/QinL233/go-plus/pkg/log"
	"github.com/QinL233/go-plus/pkg/mq/rocket"
	"github.com/QinL233/go-plus/pkg/orm/elastic"
	"github.com/QinL233/go-plus/pkg/orm/mysql"
	"github.com/QinL233/go-plus/pkg/oss/minio"
	"github.com/QinL233/go-plus/pkg/yaml"
	"github.com/gin-gonic/gin"
)

/**
服务启动入口
*/

func Start(f func(r *gin.Engine), conf ...string) {
	yaml.Init(conf...)
	log.Init()
	mysql.Init()
	elastic.Init()
	redis.Init()
	minio.Init()
	rocket.Init()
	cron.Init()
	Init(f)
}
