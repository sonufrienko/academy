package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sonufrienko/academy/api/comments/db"
	"github.com/sonufrienko/academy/api/comments/routes"
)

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func main() {
	connectString := getEnv("MONGO_URL", "")
	port := getEnv("PORT", "3000")

	err := db.Connect(connectString)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	v1 := app.Group("/v1/comments")
	v1.Get("/", routes.GetComments)
	v1.Post("/", routes.AddComments)
	v1.Delete("/:id", routes.DeleteComments)

	log.Fatal(app.Listen(":" + port))
	defer db.Disconnect()
}
