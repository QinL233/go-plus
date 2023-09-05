package rocket

import (
	"context"
	"fmt"
	"github.com/QinL233/go-plus/yaml"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"sync"
)

/**
rocketMQ consumer client
*/

var consumers []Message

type Message struct {
	//组和主题地址映射
	Group string
	Topic string
	//实际处理方法
	Handler func(msg string)
}

// Consumer 用于注册消费端
func Consumer(messages ...Message) {
	if consumers == nil {
		consumers = make([]Message, 0)
	}
	for _, message := range messages {
		consumers = append(consumers, message)
	}
}

func initConsumer(dto Message) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	newPushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{yaml.Config.Mq.RocketMq.Url}),
		consumer.WithGroupName(dto.Group))
	defer func(newPushConsumer rocketmq.PushConsumer) {
		err := newPushConsumer.Shutdown()
		if err != nil {
			log.Fatalf("rocket shutdown consumer %v err %v", dto, err)
		}
	}(newPushConsumer)

	err = newPushConsumer.Subscribe(dto.Topic, consumer.MessageSelector{},
		func(ctx context.Context, msgList ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgList {
				dto.Handler(string(msg.Body))
			}
			return consumer.ConsumeSuccess, nil
		})

	if err != nil {
		fmt.Println("rocket reader msg err!")
	}
	if err = newPushConsumer.Start(); err != nil {
		log.Fatalf("start rocket consumer %v err %v", dto, err)
	}
	fmt.Println(fmt.Sprintf("rocket connect success! %v", dto))
	wg.Wait()
}
