package infrastructure

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	router "github.com/kazuki0924/go-what-to-read-app/infrastructure/router"
)

func timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()

		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		return err
	}
}

var cacheConfig = cache.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Query("refresh") == "true"
	},
	Expiration:   15 * time.Minute,
	CacheControl: true,
}

var loggerConfig = logger.Config{
	Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
}

func SetupFiberMiddleWares() {
	router.FiberDispatcher.Use(
		timer(),
		logger.New(loggerConfig),
		cors.New(),
		cache.New(cacheConfig),
	)
}
