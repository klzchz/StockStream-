package handler

import (
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/klzchz/StockStream-/internal/market/dto"
	"github.com/klzchz/StockStream-/internal/market/entity"
	"github.com/klzchz/StockStream-/internal/market/transformer"
)

type OrderMessageHandler struct{}

func NewOrderMessageHandler() *OrderMessageHandler {
	return &OrderMessageHandler{}
}

func (h *OrderMessageHandler) ProcessOrder(msg *ckafka.Message) (*entity.Order, error) {
	fmt.Println("Received message:", string(msg.Value))
	
	var tradeInput dto.TradeInput
	if err := json.Unmarshal(msg.Value, &tradeInput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	
	order := transformer.TranformInput(tradeInput)
	return order, nil
}
