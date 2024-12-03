package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ikaushiksharma/toll-calculator/types"
)

type DataReceiver struct {
    msgch chan types.OBUData
    conn  *websocket.Conn
	Prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var(
		p DataProducer
		err error
		kafkaTopic = "obuData"
	)
	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		Prod:  p,
	}, nil
}


func (dr *DataReceiver) ProduceData(data types.OBUData) error {
	return dr.Prod.ProduceData(data)
}


func main()  {
	fmt.Println("------- Starting Data receiver --------")
	recv,_ := NewDataReceiver()
	http.HandleFunc("/ws", recv.HandleWS)
	http.ListenAndServe(":30000",nil)
}
func(dr *DataReceiver) HandleWS(w http.ResponseWriter, r *http.Request)  {
	 u:= websocket.Upgrader{
		ReadBufferSize: 1028,
		WriteBufferSize: 1028,
	 }
	conn,err := u.Upgrade(w,r, nil)
	if err!=nil{
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.wsReceiverLoop()
}
func (dr *DataReceiver) wsReceiverLoop()  {
	fmt.Println("New OBU connected Client connected")
	defer func() {
        fmt.Println("OBU Client disconnected")
        close(dr.msgch)
    }()

	
	for{
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err!= nil{
			log.Println("read error :",err)
			continue
			
		}
		if err := dr.ProduceData(data); err!= nil{
			fmt.Println("kafka produce error :", err)
		}
	}
}