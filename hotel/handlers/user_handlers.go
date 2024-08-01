package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/hotel/handlers/middlewares"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandlers struct {
	store *storage.Store
}

func NewUserHandlers(store *storage.Store) *UserHandlers {
	return &UserHandlers{
		store: store,
	}
}

// handler use to create user
// input types.CreateUserParam{} output types.User{}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	req := types.CreateUserParam{}
	if err := c.BodyParser(&req); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the request params")
	}
	msg := req.ValidateCreateUserParam()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromReq(req)
	result, err := u.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user success!",
		"data":    result,
	})
}

// handle User Login
// input types.UserLoginParams{}
func (u *UserHandlers) HandleUserLogin(c *fiber.Ctx) error {
	userloginReq := types.UserLoginParams{}
	if err := c.BodyParser(&userloginReq); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the request params")
	}
	params := userloginReq.EncryptedPassword()
	user, err := u.store.User.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	if params.Password != user.EncryptedPassword {
		return ErrorMessage(http.StatusBadRequest, "please check the email or password")
	}
	token, err := middlewares.GenToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"token":  token,
		"user":   user,
	})
}

func (u *UserHandlers) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Query("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "query param is invalid")
	}
	user, err := u.store.User.GetUserByID(c.Context(), uid)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get user success!",
		"user":    user,
	})
}

func (u *UserHandlers) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "query param is invalid")
	}
	if err := u.store.User.DeleteUserByID(c.Context(), uid); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "delete user success!",
		"status":  http.StatusOK,
	})
}

func (u *UserHandlers) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the uid param")
	}
	up := types.UpdateUserParams{}
	if err := c.BodyParser(&up); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the update params")
	}
	msg := up.ValidateUpdateUserParams()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user, err := u.store.User.UpdateUser(c.Context(), uid, &up)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserHandlers) HandleUserVerifyPassword(c *fiber.Ctx) error {
	userInfo, ok := c.Context().UserValue(middlewares.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "please login")
	}
	_ = userInfo
	return nil
}

func (u *UserHandlers) HandleUserBookingRoom(c *fiber.Ctx) error { return nil }
