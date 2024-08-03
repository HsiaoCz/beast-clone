package handlers

import (
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
)

type AdminHandlers struct {
	store *storage.Store
}

func NewAdminHandlers(store *storage.Store) *AdminHandlers {
	return &AdminHandlers{
		store: store,
	}
}

func (a *AdminHandlers) HandleCreateHotel(c *fiber.Ctx) error {
	return nil
}
