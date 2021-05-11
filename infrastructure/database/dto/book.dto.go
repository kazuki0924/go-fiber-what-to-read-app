package dto

import (
	"github.com/kazuki0924/go-what-to-read-app/domain/model"
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string      `json:"title"`
	Aka   null.String `json:"aka"`
	// Author           null.String `json:"author"`
	// PublishedAt      null.Time   `json:"published_at"`
	// Publisher        null.String `json:"publisher"`
	// TotalPages       null.Int    `json:"total_pages"`
	// ImageUrl         null.String `json:"image_url"`
	// AmazonProductUrl null.String `json:"amazon_product_url"`
	// Emoji            null.String `json:"emoji"`
}

func BookModelToBookDBModel(b *model.Book) *Book {
	book := new(Book)

	if b.Aka != "" {
		book.Aka.String = b.Aka
		book.Aka.Valid = true
	}

	book.Title = b.Title

	return book
}

// func BookDBModelToBookModel(b *Book) *model.Book {
// 	return &model.Book{
// 		Title: b.Title,
// 	}
// }
