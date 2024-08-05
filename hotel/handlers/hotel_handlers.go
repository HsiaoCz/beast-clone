package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type HotelHandlers struct {
	store *storage.Store
}

func NewHotelHandlers(store *storage.Store) *HotelHandlers {
	return &HotelHandlers{
		store: store,
	}
}

func (h *HotelHandlers) HandleGetHotels(c *fiber.Ctx) error {
	// get hotel and room don't need login
	hotels, err := h.store.Hotel.GetHotels(c.Context(), bson.M{})
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(hotels)
}
