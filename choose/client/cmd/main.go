package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/beast-clone/choose/client"
)

func main() {
	client := client.NewClient("http://localhost:9001/price")
	responsePrice, err := client.FetchPrice(context.Background(), "GG")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response %v\n", responsePrice)
}
