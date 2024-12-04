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
	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	makeHTTPTransport(*listenAddr, svc)
}
func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Printf("HTTP Transport running at port %s...\n", listenAddr)
	http.HandleFunc("/aggregate", HandleAggregate(svc))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
func HandleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&Distance); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			fmt.Fprintf(w, "Error decoding JSON: %v", err)
			return
		}
		if err := svc.AggregateDistance(Distance); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func WriteJSON(w http.ResponseWriter, httpStatus int, v any) error {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
