package rocket

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"go-plus/yaml"
)

//初始化mq
func InitMq() {
	rlog.SetLogLevel(yaml.Config.Mq.RocketMq.LogLevel)
	for _, message := range consumers {
		go initConsumer(message)
	}
}
