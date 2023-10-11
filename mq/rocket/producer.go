package rocket

import (
	"context"
	"fmt"
	"github.com/QinL233/go-plus/yaml"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"time"
)

/**
rocketMQ producer client
*/

func Producer(group, topic string, msg []byte, retry int) {
	// 消息消费失败重试两次
	newProducer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{yaml.Config.Mq.RocketMq.Url}),
		producer.WithGroupName(group),
		producer.WithRetry(retry),
	)
	defer func(newProducer rocketmq.Producer) {
		if err = newProducer.Shutdown(); err != nil {
			log.Fatalf("rocket shutdown producer err %v", err)
		}
	}(newProducer)
	if err != nil {
		log.Fatalf("rocket shutdown producer err %v", err)
	}
	if err = newProducer.Start(); err != nil {
		log.Fatalf("start shutdown producer err %v", err)
	}
	res, err := newProducer.SendSync(context.Background(), primitive.NewMessage(topic, msg))
	if err != nil {
		log.Fatalf("send msg producer err %v", err)
	}
	fmt.Printf("%s - send success : %s \n", time.Now().Format("2006-01-02 15:04:05"), res.String())
}
