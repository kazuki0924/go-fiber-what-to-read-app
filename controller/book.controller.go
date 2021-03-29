package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	"github.com/kazuki0924/go-what-to-read-app/service"
)

var (
	bookService service.BookService
)

type BookController interface {
	GetBook(c *fiber.Ctx) error
	ListBook(c *fiber.Ctx) error
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

	err = bookService.Create(book)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	c.JSON(book)
	return nil
}

func (*bookController) GetBook(c *fiber.Ctx) error {
	book := new(model.Book)

	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	book, err = bookService.Get(uint(id))
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}
	c.JSON(&book)
	return nil
}

func (*bookController) ListBook(c *fiber.Ctx) error {
	books, err := bookService.List()
	if err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}
	c.JSON(&books)
	return nil
}
