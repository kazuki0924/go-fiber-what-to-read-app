package service

import (
	repository "github.com/kazuki0924/go-what-to-read-app/domain/interface/repository"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
)

type BookService interface {
	// Validate(book *model.Book) error
	// Get(id string) (*model.Book, error)
	// List() ([]model.Book, error)
	Create(book *model.Book) error
	// Create(book *model.Book) error
	// Update(book *model.Book) error
	// Delete(book *model.Book) error
}

type bookService struct{}

var (
	bookRepository repository.BookRepository
)

func NewBookService(repository repository.BookRepository) BookService {
	bookRepository = repository
	return &bookService{}
}

func (*bookService) Create(book *model.Book) error {
	err := bookRepository.Create(book)
	if err != nil {
		return err
	}
	return nil
}
