package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/handlers/middlewares"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/gofiber/fiber/v2"
)

type RoomHandlers struct {
	store *storage.Store
}

func NewRoomHandlers(store *storage.Store) *RoomHandlers {
	return &RoomHandlers{
		store: store,
	}
}

func (rh *RoomHandlers) HandleCreateRoom(c *fiber.Ctx) error {
	// only admin user can create room
	// verify user
	userInfo, ok := c.UserContext().Value(middlewares.CtxUserInfoKey).(*types.User)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "please login")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "your have no rights to do this shit")
	}
	var create_room_params types.CreateRoomParams
	if err := c.BodyParser(&create_room_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return nil
}
