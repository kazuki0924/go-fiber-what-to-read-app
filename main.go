package main

import (
	"fmt"
	"os"

	"github.com/kazuki0924/go-what-to-read-app/controller"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	_ "github.com/kazuki0924/go-what-to-read-app/infrastructure/config"
	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
	repository "github.com/kazuki0924/go-what-to-read-app/infrastructure/repository"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
	"github.com/kazuki0924/go-what-to-read-app/service"
	"gorm.io/gorm"
)

var (
	httpRouter     router.Router = router.NewFiberRouter()
	db                           = dbFunc.InitRDB()
	dbFunc         rdb.RDB       = rdb.NewRDB()
	bookRepository               = repository.NewBookRepository(db)
	bookService                  = service.NewBookService(bookRepository)
	bookController               = controller.NewBookController(bookService)
)

func init() {
	// initialize relational database
	SetupMigrations(db)
}

// Boilerplate: add new models here
func SetupMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.Book{},
	)

	fmt.Println("Database Migrated")
}

// Boilerplate: add new routes here
func SetupRoutes(r router.Router) {
	r.POST_V1("book", bookController.CreateBook)
}

func main() {
	defer dbFunc.CloseRDB()

	// setup http routes
	SetupRoutes(httpRouter)

	// setup middlewares
	app := httpRouter.DISPATCH()
	middleware.SetupFiberMiddleWares(app)

	// listern on port $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port, app)
}
