package service

import (
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/klzchz/StockStream-/internal/market/entity"
)

type TradingService struct {
	consumer        MessageConsumer
	orderProcessor  OrderProcessor
	orderPublisher  OrderPublisher
	book           *entity.Book
	ordersIn       chan *entity.Order
	ordersOut      chan *entity.Order
	kafkaMsgChan   chan *ckafka.Message
	wg             *sync.WaitGroup
}

func NewTradingService(
	consumer MessageConsumer,
	orderProcessor OrderProcessor,
	orderPublisher OrderPublisher,
) *TradingService {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	kafkaMsgChan := make(chan *ckafka.Message)
	wg := &sync.WaitGroup{}

	book := entity.NewBook(ordersIn, ordersOut, wg)

	return &TradingService{
		consumer:       consumer,
		orderProcessor: orderProcessor,
		orderPublisher: orderPublisher,
		book:          book,
		ordersIn:      ordersIn,
		ordersOut:     ordersOut,
		kafkaMsgChan:  kafkaMsgChan,
		wg:            wg,
	}
}

func (ts *TradingService) Start() error {
	// Start consuming messages
	go ts.consumer.Consume(ts.kafkaMsgChan)

	// Start trading engine
	go ts.book.Trade()

	// Start message processing
	go ts.processMessages()

	// Start result publishing
	go ts.publishResults()

	// Wait for completion
	ts.wg.Wait()
	return nil
}

func (ts *TradingService) processMessages() {
	for msg := range ts.kafkaMsgChan {
		ts.wg.Add(1)
		order, err := ts.orderProcessor.ProcessOrder(msg)
		if err != nil {
			panic(err) // Mantendo comportamento original
		}
		ts.ordersIn <- order
	}
}

func (ts *TradingService) publishResults() {
	for order := range ts.ordersOut {
		if err := ts.orderPublisher.PublishResult(order); err != nil {
			panic(err) // Mantendo comportamento original
		}
	}
}
