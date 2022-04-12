package routes

import (
	"github.com/gorilla/mux"
	"github.com/hasnatsaeed/go-bookstore/pkg/handlers/book"
)

var RegisterBookStoreRoutes = func(router *mux.Router, handler *book.Handler) {

	router.HandleFunc("/book/", handler.Create).Methods("POST")
	router.HandleFunc("/book/", handler.FetchAll).Methods("GET")
	router.HandleFunc("/book/{bookId}", handler.FetchById).Methods("GET")
	router.HandleFunc("/book/{bookId}", handler.Update).Methods("PUT")
	router.HandleFunc("/book/{bookId}", handler.Delete).Methods("DELETE")
}
