package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/ikaushiksharma/toll-calculator/types"
)

type DataReceiver struct {
	msg  chan types.OBUData
	conn *websocket.Conn
}

func main() {
	fmt.Println("Data receiver working fine")
}