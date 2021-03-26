package infrastructure

import (
	repository "github.com/kazuki0924/go-what-to-read-app/domain/interface/repository"
	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type bookRepository struct{}

func NewBookRepository(_db *gorm.DB) repository.BookRepository {
	db = _db
	return &bookRepository{}
}

func (*bookRepository) Create(book *model.Book) error {
	err := db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}
