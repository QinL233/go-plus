package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-plus/yaml"
	"log"
	"time"
)

var driver *redis.Pool

func Init() {
	config := yaml.Config.Cache.Redis
	if config.Host == "" || config.Port == 0 {
		log.Println("redis host and port is empty")
		return
	}
	driver = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port),
				redis.DialDatabase(config.Database),
				redis.DialPassword(config.Password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	log.Println("redis connect success!")
}
