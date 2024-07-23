package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
	"github.com/HsiaoCz/beast-clone/m0NESY/db"
	"github.com/HsiaoCz/beast-clone/m0NESY/handlers"
	"github.com/HsiaoCz/beast-clone/m0NESY/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// init env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	// init db
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	// init logger
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	var (
		port         = os.Getenv("PORT")
		userDataMod  = dao.NewUserDataMod(db.Get())
		userHandlers = handlers.NewUserHandlers(userDataMod)
		router       = http.NewServeMux()
	)

	{
		// router
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))
	}

	zap.L().Info("http server is running", zap.String("listen address", port))

	http.ListenAndServe(port, router)
}
