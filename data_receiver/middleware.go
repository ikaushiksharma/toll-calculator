package main

import (
	"time"

	"github.com/ikaushiksharma/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct{
	next DataProducer
}
func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}
func(l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time)  {
		logrus.WithFields(logrus.Fields{
			"obuID":data.OBUID,
			"latitude":data.Latitiude,
			"longitude": data.Longitude,
			"took":time.Since(start),
		}).Info("producing to Kafka")	
	}(time.Now())
	return l.next.ProduceData(data)
}
