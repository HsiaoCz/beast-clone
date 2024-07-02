package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/m0NESY/handlers"
	"github.com/HsiaoCz/beast-clone/m0NESY/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := logger.InitLogger("./info.log", "./debug.log", "./error.log", "./warn.log"); err != nil {
		log.Fatal(err)
	}
	var (
		port         = os.Getenv("PORT")
		userHandlers = &handlers.UserHandlers{}
		router       = http.NewServeMux()
	)

	{
		// router
		router.HandleFunc("GET /user/hello", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))
	}

	logger.Logger.Info("http server is running", "listen addresss", port)
	
	http.ListenAndServe(os.Getenv("PORT"), router)
}
