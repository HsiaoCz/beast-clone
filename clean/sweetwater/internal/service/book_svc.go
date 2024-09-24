package service

import "github.com/gofiber/fiber/v2"

type BookService struct{}

func (b *BookService) CreateBook(c *fiber.Ctx) error {
	return nil
}
