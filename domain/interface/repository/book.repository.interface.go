package domain

import (
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
)

type BookRepository interface {
	Get(id uint) (*model.Book, error)
	List() ([]model.Book, error)
	Create(book *model.Book) error
	// Create(book *model.Book) error
	// Update(book *model.Book) error
	// Delete(book *model.Book) error
}
