package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/ikaushiksharma/toll-calculator/types"
	"google.golang.org/grpc"
)
func main() {
	httpListenAddr := flag.String("httplistenaddr", ":3000", "server listen address of http transport server")
	grpcListenAddr := flag.String("grpclistenaddr", ":8080", "server listen address of http transport server")
	flag.Parse()
	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	go makeGRPCTransport(*grpcListenAddr,svc)
	makeHTTPTransport(*httpListenAddr, svc)
}
func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Printf("HTTP Transport running at port %s...\n", listenAddr)
	http.HandleFunc("/aggregate", HandleAggregate(svc))
	http.HandleFunc("/invoice", HandleGetInvoice(svc))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Printf("gRPC Transport running at port %s...\n", listenAddr)
	// Make the TCP Listener
	listen, err := net.Listen("TCP", listenAddr)
	if err != nil {
		return err
	}
	defer listen.Close()
	// Create a new gRPC native server with oprtions
	server := grpc.NewServer([]grpc.ServerOption{}...)
	// register our GRPC server implememtation to the gRPC package
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(listen)
}


func HandleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values, ok := r.URL.Query()["obu"]
		if !ok {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})
			return
		}
		fmt.Println(values[0])
		obuId, err := strconv.Atoi(values[0])
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "inavlid OBU ID "})
		}
		fmt.Println(obuId)
		invoice, err := svc.CalculateInvoice(obuId)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		WriteJSON(w,http.StatusOK,invoice)
	}
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
