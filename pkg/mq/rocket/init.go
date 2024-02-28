package rocket

import (
	"github.com/QinL233/go-plus/pkg/yaml"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func Init() {
	rlog.SetLogLevel(yaml.Config.Mq.RocketMq.LogLevel)
	for _, message := range consumers {
		go initConsumer(message)
	}
}
