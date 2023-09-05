package rocket

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"go-plus/yaml"
)

func Init() {
	rlog.SetLogLevel(yaml.Config.Mq.RocketMq.LogLevel)
	for _, message := range consumers {
		go initConsumer(message)
	}
}
