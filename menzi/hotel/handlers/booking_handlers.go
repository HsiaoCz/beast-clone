package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/handlers/middlewares"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/HsiaoCz/beast-clone/hotel/types"
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

func (b *BookingHandlers) HandleGetBookings(c *fiber.Ctx) error {
	userInfo, ok := c.UserContext().Value(middlewares.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	_ = userInfo
	return nil
}
