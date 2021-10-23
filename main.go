package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers"
	"./models"
	"./service"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book

var db *sql.DB

// ======== Init ========

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ======== Main ========

func main() {

	db = service.ConnectDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()

	// ======== Temp-Data ========

	// books = append(books, Book{ID: 1, Title: "Les Misérables", Author: "Victor Hugo", Year: 1862},
	// 	Book{ID: 2, Title: "I Promessi sposi", Author: "Alessandro Manzoni", Year: 1847},
	// 	Book{ID: 3, Title: "Les Blancs et les Bleus", Author: "Alexandre Dumas", Year: 1867},
	// 	Book{ID: 4, Title: "A Moveable Feast", Author: "Ernest Hemingway", Year: 1964},
	// 	Book{ID: 5, Title: "Drei Kameraden", Author: "Erich Maria Remarque", Year: 1936})

	// ======== Http-Actions ========

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	// ======== Server ========

	log.Fatal(http.ListenAndServe(":8000", router))
}
