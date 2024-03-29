package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "full-cycle-kafka-kafka-1",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group2",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}