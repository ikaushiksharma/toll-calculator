package main

import (
	"log"

	"github.com/ikaushiksharma/toll-calculator/aggregator/client"
)
const (
	topic              = "obuData"
	aggregatorEndpoint = "http://localhost:3000"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	// grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	KafkaConsumer, err := NewKafkaConsumer(topic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}