package handlers

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/HsiaoCz/beast-clone/hotel/types"
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
	req := types.CreateUserParam{}
	if err := c.BodyParser(&req); err != nil {
		return NewAPIError(http.StatusBadRequest, "please check the request params")
	}
	msg := req.ValidateCreateUserParam()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user, err := types.NewUserFromReq(req)
	if err != nil {
		return NewAPIError(http.StatusInternalServerError, "create user failed please check the request params")
	}
	result, err := u.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return NewAPIError(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user success!",
		"data":    result,
	})
}

func (u *UserHandlers) HandleUserLogin(c *fiber.Ctx) error {
	userloginReq := types.UserLoginParams{}
	if err := c.BodyParser(&userloginReq); err != nil {
		return NewAPIError(http.StatusBadRequest, "please check the request params")
	}
	params, err := userloginReq.EncryptedPassword()
	if err != nil {
		return NewAPIError(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("%+v", params)
	user, err := u.store.User.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("%+v", user)
	if params.Password != user.EncryptedPassword {
		return NewAPIError(http.StatusBadRequest, "please check the email or password")
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "login success!",
	})
}
