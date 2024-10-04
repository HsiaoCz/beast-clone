package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/handlers/middlewares"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/HsiaoCz/beast-clone/hotel/types"
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
	userInfo, ok := c.UserContext().Value(middlewares.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlongin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this")
	}
	return nil
}
