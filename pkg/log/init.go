package log

import (
	"log"
)

func Init() {
	//设置日志格式
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//初始化logstash
	initLogstash()
}
