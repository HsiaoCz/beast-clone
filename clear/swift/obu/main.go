package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/HsiaoCz/beast-clone/swift/types"
	"github.com/gorilla/websocket"
)

const wsEnpoint = "ws://127.0.0.1:30001/ws"
const sendInterval = time.Second

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {
	obuIDS := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEnpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%+v\n", data)
		}
		time.Sleep(sendInterval)
	}
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
