package main

import (
	"log"
)
const topic = "obuData"

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	KafkaConsumer, err := NewKafkaConsumer(topic, svc)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}