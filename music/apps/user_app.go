package apps

import "github.com/gofiber/fiber/v2"

type UserApp struct{}

func UserAppInit() *UserApp {
	return &UserApp{}
}

func (u *UserApp) HandleUserCreate(c *fiber.Ctx) error {
	return nil
}
