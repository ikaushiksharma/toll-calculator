package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ikaushiksharma/toll-calculator/types"
)
func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server listen address of http transport server")
	flag.Parse()
	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)
	makeHTTPTransport(*listenAddr,svc)
}
func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Printf("HTTP Transport running at port %s", listenAddr)
	http.HandleFunc("/aggregate", HandleAggregate(svc))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
func HandleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&Distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding JSON: %v", err)
			return
		}
	}
}
