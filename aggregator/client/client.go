package client

import (
	"context"

	"github.com/ikaushiksharma/toll-calculator/types"
)
type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error 		//being used by distance calculator to aggregate the tolls
	GetInvoice(context.Context, int) (*types.Invoice, error)  		//being used by gateway to get the invoice
}
