package main

import (
	"fmt"
	"log"

	"github.com/HsiaoCz/beast-clone/gustao/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("all is well")
}
