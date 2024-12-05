package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ikaushiksharma/toll-calculator/types"
)
type HTTPClient struct {
	EndPoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		EndPoint: endpoint,
	}
}

// the client here is putting all the distance data coming from distance calculator to the aggregator service
func (c *HTTPClient) AggregateInvoice(distance types.Distance) error {
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.EndPoint, bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("The service responded with a non 200 status code %d", resp.StatusCode)
	}
	return nil
}
