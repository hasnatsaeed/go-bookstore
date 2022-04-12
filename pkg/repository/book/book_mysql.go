package book

import (
	"github.com/hasnatsaeed/go-bookstore/pkg/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func NewMySqlBookRepository(db *gorm.DB) Repository {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return &mysqlPostRepo{
		db: db,
	}
}

type mysqlPostRepo struct {
	db *gorm.DB
}

func (repo *mysqlPostRepo) CreateBook(b *models.Book) (*models.Book, error) {
	repo.db.Create(&b)
	return b, nil
}

func (repo *mysqlPostRepo) FetchAllBooks() (*[]models.Book, error) {

	books := make([]models.Book, 50)

	repo.db.Find(&books)

	return &books, nil
}

func (repo *mysqlPostRepo) FetchBookById(id int64) (*models.Book, error) {

	book := &models.Book{}

	repo.db.Where("ID=?", id).Find(book)

	return book, nil
}

func (repo *mysqlPostRepo) DeleteBook(id int64) (*models.Book, error) {

	book := &models.Book{}
	repo.db.Where("ID=?", id).Delete(book)
	return book, nil
}

func (repo *mysqlPostRepo) UpdateBook(bookId int64, update *models.Book) (*models.Book, error) {
	bookToUpdate, _ := repo.FetchBookById(bookId)

	if update.Name != "" {
		bookToUpdate.Name = update.Name
	}
	if update.Author != "" {
		bookToUpdate.Author = update.Author
	}
	if update.Publication != "" {
		bookToUpdate.Publication = update.Publication
	}

	repo.db.Save(bookToUpdate)

	return bookToUpdate, nil
}
