package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	database "github.com/kazuki0924/go-what-to-read-app/infrastructure/database"
	env "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)
	err := c.BodyParser(book)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	db.Create(&book)
	c.JSON(book)
	return nil
}

func main() {
	env.LoadEnv()

	port := os.Getenv("HTTP_PORT")

	app := fiber.New()

	app.Post("api/v1/book", CreateBook)

	database.InitDB()
	defer database.CloseDB()

	app.Listen(":" + port)
}
