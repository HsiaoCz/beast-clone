package handlers

import (
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
)

type HotelHandlers struct {
	store *storage.Store
}

func NewHotelHandlers(store *storage.Store) *HotelHandlers {
	return &HotelHandlers{
		store: store,
	}
}

func (h *HotelHandlers)HandleGetHotels(c *fiber.Ctx)error{
	return nil
}