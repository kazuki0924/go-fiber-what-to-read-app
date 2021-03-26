package infrastructure

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type fiberRouter struct {
}

var (
	FiberDispatcher = fiber.New()
	api             = FiberDispatcher.Group("/api")
	v1              = api.Group("/v1")
)

type handler = func(c *fiber.Ctx) error

func NewFiberRouter() Router {
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
	fmt.Printf("Fiber HTTP server running on port %v", port)
	FiberDispatcher.Listen(":" + port)
}
