package client

import (
	"context"

	"github.com/ikaushiksharma/toll-calculator/types"
)
type Client interface{
	Aggregate(context.Context, *types.AggregateRequest) (error)
}
