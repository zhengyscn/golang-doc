package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

func Init(brokers []string) (err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Send(lineStr string) {
	msg := &sarama.ProducerMessage{
		Topic: "important",
		Value: sarama.StringEncoder(lineStr),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Message is stored in partition(%d)/offset(%d)\n", partition, offset)
}

func Kclosed() {
	producer.Close()
}
