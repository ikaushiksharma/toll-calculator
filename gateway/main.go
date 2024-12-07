package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ikaushiksharma/toll-calculator/aggregator/client"
	"github.com/sirupsen/logrus"
)
type apifunc func(w http.ResponseWriter, r *http.Request) error

type InvoiceHandler struct {
	Client client.Client
}

func NewInvoiceHandler(c client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		Client: c,
	}
}

func makeApifunc(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func (start time.Time)  {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri": r.RequestURI,
			}).Info("REQ::")
		}(time.Now())

		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "HTTP server listen address")
	flag.Parse()
	aggregatorServiceAddr := flag.String("aggregatorServiceAddr", "http://localhost:4000", "aggregator server listen address")
	flag.Parse()
	var (
		client     = client.NewHTTPClient(*aggregatorServiceAddr) // endpoint of the aggregator service
		invHandler = NewInvoiceHandler(client)
	)
	http.HandleFunc("/invoice", makeApifunc(invHandler.HandleGetInvoice))
	logrus.Infof("HTTP Gateway server is running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func (h *InvoiceHandler) HandleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("hitting the invoice endpoint inside gateway......")
	// access agg client
	inv, err := h.Client.GetInvoice(r.Context(), 1373793167)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, inv)
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}