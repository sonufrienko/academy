package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sonufrienko/academy/api/comments/routes"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1/comments")
	v1.Get("/", routes.GetComments)
	v1.Post("/", routes.AddComments)
	v1.Delete("/:id", routes.DeleteComments)

	log.Fatal(app.Listen(":3000"))
}
