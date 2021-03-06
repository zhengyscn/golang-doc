package main

import (
	"fmt"

	"github.com/zhengyscn/golang-doc/kafka/kafka"
	"github.com/zhengyscn/golang-doc/kafka/tailf"
)

func main() {
	var (
		err      error
		topic    string   = "important"
		filename string   = "/tmp/nginx.log"
		brokers  []string = []string{"localhost:9092"}
	)

	err = kafka.InitProducer(brokers)
	if err != nil {
		fmt.Sprintf("init kafka err: %v", err)
	}
	defer kafka.ProducerClose()
	fmt.Println("init kafka succ.")

	err = tailf.Init(filename)
	if err != nil {
		fmt.Sprintf("init tailf err: %v", err)
	}
	fmt.Println("init tail succ.")

	var msgChan = make(chan string)
	go tailf.ReadLine(msgChan)
	for {
		select {
		case msg := <-msgChan:
			kafka.Send(msg, topic)
		}
	}

}
