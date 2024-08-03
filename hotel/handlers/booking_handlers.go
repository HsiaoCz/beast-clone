package handlers

import (
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
)

type BookingHandlers struct {
	store *storage.Store
}

func NewBookingHandlers(store *storage.Store) *BookingHandlers {
	return &BookingHandlers{
		store: store,
	}
}

func (b *BookingHandlers)HandleGetBookings(c *fiber.Ctx)error{
	return nil
}