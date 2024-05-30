package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	store *storage.Store
}

func NewUserHandlers(store *storage.Store) *UserHandlers {
	return &UserHandlers{
		store: store,
	}
}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user success!",
	})
}
