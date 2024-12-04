package main

import (
	"log"

	"github.com/ikaushiksharma/toll-calculator/aggregator/client"
)
const (
	topic              = "obuData"
	aggregatorEndpoint = "http://localhost:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	KafkaConsumer, err := NewKafkaConsumer(topic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}