package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ikaushiksharma/toll-calculator/types"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"
var sendInterval = time.Second
func genCord() float64 {
	n := float64(rand.Intn(100)+1)
	f := rand.Float64()
	return n + f
}
func Location() (float64,float64){
	return genCord(),genCord()
}
func generateOBUID(n int) []int {
	ids := make([]int,n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}
func main()  {
	obuIDs := generateOBUID(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint,nil)
	if err!= nil{
		log.Fatal(err)
	}
	for{
		for i := 0; i < len(obuIDs); i++ {
			lat,long := Location()
			data := types.OBUData{
				OBUID: obuIDs[i],
				Latitiude: lat,
				Longitude: long,
			}
			fmt.Printf("%+v\n",data)
			if err := conn.WriteJSON(data); err!= nil{
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}