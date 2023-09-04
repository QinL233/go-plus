package app

import (
	"go-plus/gin"
	"go-plus/yaml"
)

func Start(conf ...string) {
	yaml.Init(conf...)
	gin.Init()
}
