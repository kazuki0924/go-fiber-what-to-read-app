package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	env "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
	"gorm.io/gorm"

	repository "github.com/kazuki0924/go-what-to-read-app/infrastructure/repository"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
)

var (
	dbFunc     rdb.RDB       = rdb.NewRDB()
	httpRouter router.Router = router.NewFiberRouter()
	db         *gorm.DB
)

func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	err := c.BodyParser(book)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	var repo = repository.NewBookRepository(db)
	repo.Create(book)
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
	_db, err := dbFunc.InitRDB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db = _db
	defer dbFunc.CloseRDB()

	// setup http routes
	SetupRoutes()

	// setup middlewares
	middleware.SetupFiberMiddleWares()

	// listern on port: $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port)
}
