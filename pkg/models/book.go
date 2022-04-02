package models

import (
	"github.com/hasnatsaeed/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {

	var books []Book

	db.Find(&books)

	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {

	var book Book

	db := db.Where("ID=?", id).Find(&book)

	return &book, db

}

func DeleteBook(id int64) *Book {

	var book *Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
