package service

import (
	"github.com/HsiaoCz/beast-clone/clean/sweetwater/internal/data"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	dau data.UserDataInter
}

func UserServiceInit(dau data.UserDataInter) *UserService {
	return &UserService{
		dau: dau,
	}
}

func (u *UserService) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (u *UserService) DeletedUser(c *fiber.Ctx) error {
	return nil
}
