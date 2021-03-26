package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	infrastructure "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
)

func main() {
	infrastructure.LoadEnv()

	port := os.Getenv("HTTP_PORT")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + port)

}
