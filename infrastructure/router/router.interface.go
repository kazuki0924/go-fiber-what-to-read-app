package infrastructure

import "github.com/gofiber/fiber/v2"

type Router interface {
	GET_V1(uri string, f func(c *fiber.Ctx) error)
	POST_V1(uri string, f func(c *fiber.Ctx) error)
	PUT_V1(uri string, f func(c *fiber.Ctx) error)
	DELETE_V1(uri string, f func(c *fiber.Ctx) error)
	SERVE(port string)
}
