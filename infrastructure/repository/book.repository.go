package infrastructure

import (
	"errors"

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

func (*bookRepository) Get(id uint) (*model.Book, error) {
	var book *model.Book
	err := db.Find(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return book, nil
}
