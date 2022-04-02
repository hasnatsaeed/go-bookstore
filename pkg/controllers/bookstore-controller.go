package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hasnatsaeed/go-bookstore/pkg/models"
	"github.com/hasnatsaeed/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	books := models.GetAllBooks()

	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing 'bookId'")
	}
	book, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing 'bookId'")
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing 'bookId'")
	}

	book := &models.Book{}
	utils.ParseBody(r, book)

	bookToUpdate, db := models.GetBookById(bookId)

	if book.Name != "" {
		bookToUpdate.Name = book.Name
	}
	if book.Author != "" {
		bookToUpdate.Author = book.Author
	}
	if book.Publication != "" {
		bookToUpdate.Publication = book.Publication
	}

	db.Save(bookToUpdate)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
