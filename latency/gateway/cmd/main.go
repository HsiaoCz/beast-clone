package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/latency/gateway/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(handlers.ErrorMsg); ok {
			return c.Status(e.Status).JSON(&e)
		}
		errMsg := handlers.ErrorMsg{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(errMsg.Status).JSON(&errMsg)
	},
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		port         = os.Getenv("PORT")
		userHandlers = &handlers.UserHandlers{}
		router       = fiber.New(config)
		v1           = router.Group("/api/v1")
	)
	{
		// user router
		v1.Post("/user", userHandlers.HandleCreateUser)
		v1.Delete("/user/:uid", userHandlers.HandleDeleteUser)
		v1.Get("/user/:uid", userHandlers.HandleGetUserByID)
		v1.Get("/user", userHandlers.HandleGetUsers)
	}
	router.Listen(port)

}
