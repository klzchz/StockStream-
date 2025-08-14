package service

import (
	"github.com/klzchz/StockStream-/internal/infra/kafka"
)

type KafkaProducerAdapter struct {
	producer *kafka.Producer
}

func NewKafkaProducerAdapter(producer *kafka.Producer) *KafkaProducerAdapter {
	return &KafkaProducerAdapter{
		producer: producer,
	}
}

func (k *KafkaProducerAdapter) Publish(msg interface{}, key []byte, topic string) error {
	return k.producer.Publish(msg, key, topic)
}
