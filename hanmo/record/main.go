package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/api/v1/record", HandleCreateRecord)

	app.Listen(":3001")
}
