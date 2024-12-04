package main

import (
	"fmt"

	"github.com/ikaushiksharma/toll-calculator/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}
type Storer interface {
	Insert(types.Distance) error
}
type InvoiceAggregator struct {
	store Storer
}
func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("Processing and Inserting Distance in the storage...", distance)
	return i.store.Insert(distance)
}
func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}
