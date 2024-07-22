package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3001", "listen address the server is running")
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
