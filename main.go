package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book ...
type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var book []Book

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{ID}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/book/{ID}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// ======== Book-Actions ========

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets a books")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("New book is added")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("The book is updated")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("The book is deleted")
}
