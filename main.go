package main

import (
	"log"
	"os"

	env "github.com/kazuki0924/go-what-to-read-app/config/env"
	"github.com/kazuki0924/go-what-to-read-app/controller"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
	repository "github.com/kazuki0924/go-what-to-read-app/infrastructure/repository"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
	"github.com/kazuki0924/go-what-to-read-app/service"
	"gorm.io/gorm"
)

var (
	httpRouter     router.Router = router.NewFiberRouter()
	dbFunc         rdb.RDB       = rdb.NewRDB()
	db             *gorm.DB
	bookRepository = repository.NewBookRepository(db)
	bookService    = service.NewBookService(bookRepository)
	bookController = controller.NewBookController(bookService)
)

// Boilerplate: add new routes here
func SetupRoutes(r router.Router) {
	r.POST_V1("book", bookController.CreateBook)
}

// Boilerplate: add new models here
func SetupMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.Book{},
	)
}

func main() {
	// load environment variables
	env.LoadEnv()

	// initialize relational database
	var err error
	db, err = dbFunc.InitRDB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer dbFunc.CloseRDB()

	// setup http routes
	SetupRoutes(httpRouter)

	// setup middlewares
	app := httpRouter.DISPATCH()
	middleware.SetupFiberMiddleWares(app)

	// listern on port $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port)
}
