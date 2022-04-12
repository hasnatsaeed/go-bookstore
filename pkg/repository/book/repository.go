package book

import (
	"github.com/hasnatsaeed/go-bookstore/pkg/models"
)

type Repository interface {
	CreateBook(b *models.Book) (*models.Book, error)
	UpdateBook(id int64, b *models.Book) (*models.Book, error)
	FetchBookById(id int64) (*models.Book, error)
	FetchAllBooks() (*[]models.Book, error)
	DeleteBook(id int64) (*models.Book, error)
}
