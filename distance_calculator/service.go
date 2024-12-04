package main

import (
	"fmt"
	"math"

	"github.com/ikaushiksharma/toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct {
	PrevPoint []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("calculating the distance")
	distance := 0.0

	if len(s.PrevPoint) > 0 {
		distance = CalculateDistance(s.PrevPoint[0], data.Latitiude, s.PrevPoint[1], data.Longitude)
	}
	s.PrevPoint = []float64{data.Latitiude, data.Longitude}
	return distance, nil
}

func CalculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}