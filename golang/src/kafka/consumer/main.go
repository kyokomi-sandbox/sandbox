package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":  "FirstAppConsumerGroup",
		"enable.auto.commit": "false",
	})
	defer c.Close()

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{"first-app"}, nil); err != nil {
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			c.CommitMessage(msg)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
