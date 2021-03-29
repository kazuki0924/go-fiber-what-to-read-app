package infrastructure

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	middleware "github.com/kazuki0924/go-what-to-read-app/infrastructure/middleware"
)

type fiberRouter struct{}

var (
	FiberDispatcher *fiber.App
	app             fiber.Router
	api             fiber.Router
	v1              fiber.Router
)

type handler = func(c *fiber.Ctx) error

func NewFiberRouterWithMiddlewares() Router {
	FiberDispatcher = fiber.New()
	app = middleware.SetupFiberMiddleWares(FiberDispatcher)
	api = app.Group("/api")
	v1 = api.Group("/v1")
	return &fiberRouter{}
}

func (*fiberRouter) GET_V1(uri string, f handler) {
	v1.Get(uri, f)
}

func (*fiberRouter) POST_V1(uri string, f handler) {
	v1.Post(uri, f)
}

func (*fiberRouter) PUT_V1(uri string, f handler) {
	v1.Put(uri, f)
}

func (*fiberRouter) DELETE_V1(uri string, f handler) {
	v1.Delete(uri, f)
}

func (*fiberRouter) SERVE(port string) {
	fmt.Printf("Fiber HTTP server running on port %v\n", port)

	log.Fatal(FiberDispatcher.Listen(":" + port))
}
