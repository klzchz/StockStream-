package config

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConfig struct {
	BootstrapServers string
	GroupID          string
	AutoOffsetReset  string
	InputTopic       string
	OutputTopic      string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		BootstrapServers: "host.docker.internal:9094",
		GroupID:          "myGroup",
		AutoOffsetReset:  "latest",
		InputTopic:       "input",
		OutputTopic:      "output",
	}
}

func (c *KafkaConfig) ToConfigMap() *ckafka.ConfigMap {
	return &ckafka.ConfigMap{
		"bootstrap.servers": c.BootstrapServers,
		"group.id":          c.GroupID,
		"auto.offset.reset": c.AutoOffsetReset,
	}
}
