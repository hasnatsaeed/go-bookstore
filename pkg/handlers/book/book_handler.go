package book

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hasnatsaeed/go-bookstore/pkg/models"
	"github.com/hasnatsaeed/go-bookstore/pkg/repository/book"
	"github.com/hasnatsaeed/go-bookstore/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func NewBookHandler(db *gorm.DB) *Handler {
	return &Handler{
		repository: book.NewMySqlBookRepository(db),
	}
}

type Handler struct {
	repository book.Repository
}

func (book *Handler) FetchAll(w http.ResponseWriter, r *http.Request) {

	books, _ := book.repository.FetchAllBooks()

	respondWithJSON(w, http.StatusOK, books)
}

func (book *Handler) FetchById(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(mux.Vars(r)["bookId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error while parsing 'bookId'")
	}
	bookById, _ := book.repository.FetchBookById(int64(bookId))
	respondWithJSON(w, http.StatusOK, bookById)
}

func (book *Handler) Create(w http.ResponseWriter, r *http.Request) {
	bookToCreate := &models.Book{}
	utils.ParseBody(r, bookToCreate)
	createdBook, _ := book.repository.CreateBook(bookToCreate)
	respondWithJSON(w, http.StatusOK, createdBook)
}

func (book *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(mux.Vars(r)["bookId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error while parsing 'bookId'")
	}
	deletedBook, _ := book.repository.DeleteBook(int64(bookId))
	respondWithJSON(w, http.StatusOK, deletedBook)
}

func (book *Handler) Update(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.Atoi(mux.Vars(r)["bookId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error while parsing 'bookId'")
	}
	update := &models.Book{}
	utils.ParseBody(r, update)

	updatedBook, _ := book.repository.UpdateBook(int64(bookId), update)

	respondWithJSON(w, http.StatusOK, updatedBook)

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
