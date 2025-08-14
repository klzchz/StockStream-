package main

import (
	"github.com/klzchz/StockStream-/internal/application/handler"
	"github.com/klzchz/StockStream-/internal/application/service"
	"github.com/klzchz/StockStream-/internal/config"
	"github.com/klzchz/StockStream-/internal/infra/kafka"
)

func main() {
	// Configuration (SRP)
	kafkaConfig := config.NewKafkaConfig()
	configMap := kafkaConfig.ToConfigMap()

	// Infrastructure setup
	producer := kafka.NewKafkaProducer(configMap)
	consumer := kafka.NewConsumer(configMap, []string{kafkaConfig.InputTopic})

	// Adapters (DIP)
	consumerAdapter := service.NewKafkaConsumerAdapter(consumer)
	producerAdapter := service.NewKafkaProducerAdapter(producer)

	// Handlers (SRP)
	orderHandler := handler.NewOrderMessageHandler()
	resultPublisher := handler.NewResultPublisher(producerAdapter, kafkaConfig)

	// Application Service (orchestrates the flow)
	tradingService := service.NewTradingService(
		consumerAdapter,
		orderHandler,
		resultPublisher,
	)

	// Start the application
	if err := tradingService.Start(); err != nil {
		panic(err)
	}
}
