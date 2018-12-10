package kafka

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var (
	err    error
	master sarama.Consumer
)

func InitConsumer(brokers []string) (err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	master, err = sarama.NewConsumer(brokers, config)
	return

}

func Consumer(topic string) (err error) {
	// How to decide partition, is it fixed value...?
	pt, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-pt.Errors():
				fmt.Println(err)
			case msg := <-pt.Messages():
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	return
}

func ConsumerClose() {
	master.Close()
}
