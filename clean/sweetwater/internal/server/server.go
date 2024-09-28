package server

import (
	"github.com/HsiaoCz/beast-clone/clean/sweetwater/internal/data"
	"github.com/HsiaoCz/beast-clone/clean/sweetwater/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type HttpServer struct {
	app    *fiber.App
	client *mongo.Client
}

func HttpServerInit(app *fiber.App, client *mongo.Client) *HttpServer {
	return &HttpServer{
		client: client,
		app:    app,
	}
}

func (h *HttpServer) Register(port string) error {
	var (
		userData = data.UserDataInit(h.client, h.client.Database("swt").Collection("users"))
		userSvc  = service.UserServiceInit(userData)
	)

	{
		h.app.Post("/user", userSvc.CreateUser)
		h.app.Delete("/user/:uid", userSvc.DeletedUser)
	}
	return h.app.Listen(port)
}
