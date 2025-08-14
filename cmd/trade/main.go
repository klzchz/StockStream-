package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/klzchz/StockStream-/internal/infra/kafka"
	"github.com/klzchz/StockStream-/internal/market/dto"
	"github.com/klzchz/StockStream-/internal/market/entity"
	"github.com/klzchz/StockStream-/internal/market/transformer"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)

	wg := &sync.WaitGroup{}

	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan) //t2

	//recebe de kafka
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() //t3

	go func() {
		for msg := range kafkaMsgChan {
			// Process the message received from Kafka
			fmt.Println("Received message:", string(msg.Value))
			wg.Add(1)
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TranformInput(tradeInput)
			ordersIn <- order

		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			panic(err)
		}
		err = producer.Publish(outputJson, []byte("orders"), "output")
		if err != nil {
			panic(err)
		}
	}

}
