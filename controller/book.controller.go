package controller

import (
	"github.com/gofiber/fiber/v2"

	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	"github.com/kazuki0924/go-what-to-read-app/service"
)

var (
	bookService service.BookService
)

type BookController interface {
	// GetBook(c *fiber.Ctx) error
	// ListBook(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
	// UpdateBook(c *fiber.Ctx) error
	// DeleteBook(c *fiber.Ctx) error
}

type bookController struct{}

func NewBookController(service service.BookService) BookController {
	bookService = service
	return &bookController{}
}

func (*bookController) CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	err := c.BodyParser(book)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	bookService.Create(book)
	c.JSON(book)
	return nil
}
