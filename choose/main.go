package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewLoggingService(NewPriceFetcher())
	price, err := svc.FetchPrice(context.Background(), "GG")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
