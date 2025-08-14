package service

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/klzchz/StockStream-/internal/infra/kafka"
)

type KafkaConsumerAdapter struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumerAdapter(consumer *kafka.Consumer) *KafkaConsumerAdapter {
	return &KafkaConsumerAdapter{
		consumer: consumer,
	}
}

func (k *KafkaConsumerAdapter) Consume(msgChan chan *ckafka.Message) error {
	return k.consumer.Consume(msgChan)
}
