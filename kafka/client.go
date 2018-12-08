package main

import (
	"fmt"

	"github.com/zhengyscn/golang-doc/kafka/kafka"
)

func main() {
	var (
		err     error
		topic   string   = "important"
		brokers []string = []string{"localhost:9092"}
	)
	err = kafka.InitConsumer(brokers)
	if err != nil {
		fmt.Printf("init kafka consumer err: %v\n", err)
	}
	err = kafka.Consumer(topic)
	if err != nil {
		fmt.Printf("kafka consumer err: %v\n", err)
	}
	defer kafka.ConsumerClose()
}
