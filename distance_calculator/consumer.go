package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ikaushiksharma/toll-calculator/types"
	"github.com/sirupsen/logrus"
)
type KafkaConsumer struct {
	consumer    *kafka.Consumer
	IsRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println("error in creating consumer!")
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info(" Kafka Transport started")
	c.IsRunning = true
	c.ReadMessageLoop()
}

func (c *KafkaConsumer) ReadMessageLoop() {
	for c.IsRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {

			logrus.Errorf("kafka consume error: %v (%v)\n", err, msg)
		}

		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error: %v", err)
			continue
		}

		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Error in calculating distance: %v", err)
			continue
		}
		_ = distance
		fmt.Printf("distance calculated: %.2f\n ", distance)

	}
}