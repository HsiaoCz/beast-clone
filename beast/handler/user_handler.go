package handler

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/beast/data"
	"github.com/HsiaoCz/beast-clone/beast/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	user data.UserStorer
}

func NewUserHandler(userStorer data.UserStorer) *UserHandler {
	return &UserHandler{
		user: userStorer,
	}
}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var userCreateParams types.CreateUserParams
	if err := c.BodyParser(&userCreateParams); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := userCreateParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(userCreateParams)
	result, err := u.user.CreateUser(user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    result,
	})
}

func (u *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	user_id := c.Query("user_id")
	user, err := u.user.GetUserByID(user_id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "get user success",
		"user":    user,
	})
}

func (u *UserHandler) HandleUserLogin(c *fiber.Ctx) error {
	return nil
}

func (u *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	return nil
}
