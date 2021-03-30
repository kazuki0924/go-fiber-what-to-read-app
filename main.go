package main

import (
	"os"

	"github.com/kazuki0924/go-what-to-read-app/controller"
	_ "github.com/kazuki0924/go-what-to-read-app/infrastructure/config/env"
	"github.com/kazuki0924/go-what-to-read-app/infrastructure/database/dto"
	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	repository "github.com/kazuki0924/go-what-to-read-app/infrastructure/repository"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
	"github.com/kazuki0924/go-what-to-read-app/service"
)

var (
	httpRouter     router.Router = router.NewFiberRouterWithMiddlewares()
	db                           = dbFunc.InitRDB()
	dbFunc         rdb.RDB       = rdb.NewRDB()
	bookRepository               = repository.NewBookRepository(db)
	bookService                  = service.NewBookService(bookRepository)
	bookController               = controller.NewBookController(bookService)
)

// Boilerplate: add new routes here
func SetupRoutes(r router.Router) {
	r.GET_V1("book/:id", bookController.GetBook)
	r.GET_V1("books", bookController.ListBook)
	r.POST_V1("book", bookController.CreateBook)
}

func main() {
	// initialize relational database
	// Boilerplate: add new models here
	dbFunc.SetupMigrations(
		&dto.Book{},
	)

	defer dbFunc.CloseRDB()

	// setup http routes
	SetupRoutes(httpRouter)

	// listern on port $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port)
}
