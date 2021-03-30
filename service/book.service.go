package service

import (
	repository "github.com/kazuki0924/go-what-to-read-app/domain/interface/repository"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	"github.com/kazuki0924/go-what-to-read-app/infrastructure/database/dto"
)

type BookService interface {
	// Validate(book *model.Book) error
	Get(id uint) (*model.Book, error)
	List() ([]model.Book, error)
	Create(book *dto.Book) error
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

func (*bookService) Create(book *dto.Book) error {
	err := bookRepository.Create(book)
	if err != nil {
		return err
	}
	return nil
}

func (*bookService) Get(id uint) (*model.Book, error) {
	book, err := bookRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*bookService) List() ([]model.Book, error) {
	books, err := bookRepository.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}
