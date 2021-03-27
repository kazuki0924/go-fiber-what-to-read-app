package infrastructure

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	repository "github.com/kazuki0924/go-what-to-read-app/infrastructure/repository"
)

type fiberRouter struct {
}

var (
	FiberDispatcher = fiber.New()
	api             = FiberDispatcher.Group("/api")
	v1              = api.Group("/v1")
	db              *gorm.DB
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

func SetupRoutes(r Router) {
	r.POST_V1("book", CreateBook)
}
