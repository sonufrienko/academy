package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/comments", func(c *fiber.Ctx) error {
		return c.SendString("Comments goes here")
	})

	app.Listen(":3000")
}
