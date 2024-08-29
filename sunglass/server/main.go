package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/beast-clone/sunglass"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	sunglass.NewGroup("scores", 2<<10, sunglass.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	peers := sunglass.NewHTTPPool(addr)
	log.Println("sunglass cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
