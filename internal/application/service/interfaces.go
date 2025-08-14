package service

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/klzchz/StockStream-/internal/market/entity"
)

// MessageConsumer abstrai o consumo de mensagens
type MessageConsumer interface {
	Consume(msgChan chan *ckafka.Message) error
}

// MessageProducer abstrai a produção de mensagens
type MessageProducer interface {
	Publish(msg interface{}, key []byte, topic string) error
}

// OrderProcessor abstrai o processamento de ordens
type OrderProcessor interface {
	ProcessOrder(msg *ckafka.Message) (*entity.Order, error)
}

// OrderPublisher abstrai a publicação de resultados
type OrderPublisher interface {
	PublishResult(order *entity.Order) error
}
