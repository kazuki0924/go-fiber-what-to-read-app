package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	env "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
	"gorm.io/gorm"
)

var (
	conn       rdb.RDB       = rdb.NewRDB()
	httpRouter router.Router = router.NewFiberRouter()
)

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func CreateBook(c *fiber.Ctx) error {
	db := rdb.RDBConn

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

func SetupRoutes() {
	httpRouter.POST_V1("book", CreateBook)
}

func main() {
	// load environment variables
	env.LoadEnv()

	// initialize relational database
	conn.InitRDB()
	defer conn.CloseRDB()

	// setup http routes
	SetupRoutes()

	// setup middlewares
	middleware.SetupFiberMiddleWares()

	// listern on port: $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port)
}
