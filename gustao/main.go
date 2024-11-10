package main

import (
	"fmt"
	"log"

	"github.com/HsiaoCz/beast-clone/gustao/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
    
	fmt.Println("all is well")
}
