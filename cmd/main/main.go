package main

import (
	"github.com/gorilla/mux"
	"github.com/hasnatsaeed/go-bookstore/pkg/handlers/book"
	"github.com/hasnatsaeed/go-bookstore/pkg/routes"
	"github.com/hasnatsaeed/go-bookstore/pkg/storage"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	db := connectMySqlDb()

	bookHandler := book.NewBookHandler(db)

	routes.RegisterBookStoreRoutes(router, bookHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}

func connectMySqlDb() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db, err := storage.ConnectMySql(dbHost, dbPort, dbUser, dbPass, dbName)

	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return db
}
