package infrastructure

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
)

var cacheConfig = cache.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Query("refresh") == "true"
	},
	Expiration:   15 * time.Minute,
	CacheControl: true,
}

func SetupFiberMiddleWares() {
	router.FiberDispatcher.Use(
		logger.New(),
		cors.New(),
		cache.New(cacheConfig),
	)
}
