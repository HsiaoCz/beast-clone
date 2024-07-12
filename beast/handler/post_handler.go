package handler

import "github.com/gofiber/fiber/v2"

type PostHandler struct{}

func (p *PostHandler) HandleCreatePost(c *fiber.Ctx) error {
	return nil
}
