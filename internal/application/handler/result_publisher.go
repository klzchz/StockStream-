package handler

import (
	"encoding/json"
	"fmt"

	"github.com/klzchz/StockStream-/internal/application/service"
	"github.com/klzchz/StockStream-/internal/config"
	"github.com/klzchz/StockStream-/internal/market/entity"
	"github.com/klzchz/StockStream-/internal/market/transformer"
)

type ResultPublisher struct {
	producer service.MessageProducer
	config   *config.KafkaConfig
}

func NewResultPublisher(producer service.MessageProducer, config *config.KafkaConfig) *ResultPublisher {
	return &ResultPublisher{
		producer: producer,
		config:   config,
	}
}

func (r *ResultPublisher) PublishResult(order *entity.Order) error {
	output := transformer.TransformOutput(order)
	outputJson, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output: %w", err)
	}
	
	fmt.Println("Processed order:", string(outputJson))
	
	if err := r.producer.Publish(outputJson, []byte("orders"), r.config.OutputTopic); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	
	return nil
}
