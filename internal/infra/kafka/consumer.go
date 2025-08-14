package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
	consumer  *ckafka.Consumer
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) (*Consumer, error) {
	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	c := &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
		consumer:  consumer,
	}

	// Subscribe to topics
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		consumer.Close()
		return nil, fmt.Errorf("failed to subscribe to topics: %w", err)
	}

	return c, nil
}

func (c *Consumer) Consume() (*ckafka.Message, error) {
	msg, err := c.consumer.ReadMessage(-1)
	if err != nil {
		return nil, fmt.Errorf("consumer error: %w", err)
	}
	return msg, nil
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}
