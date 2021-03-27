package main

import (
	"log"
	"os"

	rdb "github.com/kazuki0924/go-what-to-read-app/infrastructure/database/rdb"
	env "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
	"gorm.io/gorm"

	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
)

var (
	dbFunc     rdb.RDB       = rdb.NewRDB()
	httpRouter router.Router = router.NewFiberRouter()
	db         *gorm.DB
)

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
	router.SetupRoutes(httpRouter)

	// setup middlewares
	middleware.SetupFiberMiddleWares()

	// listern on port: $HTTP_PORT
	port := os.Getenv("HTTP_PORT")
	httpRouter.SERVE(port)
}
