package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"../models"

	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct{}

// ======== Init

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ======== Actions

// GetBooks ...
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Works

		var book models.Book
		books = []models.Book{}

		rows, err := db.Query("select * from books")
		logFatal(err)

		// defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)

			books = append(books, book)
		}

		json.NewEncoder(w).Encode(books)

		// Status
		log.Println("Get all books")
	}
}

// -------------------------------------

// GetBook ...
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Not Working :(

		// ======== Temp Code
		// params := mux.Vars(r)
		// i, _ := strconv.Atoi(params["ID"]) // Convert String => Int
		// Info Type | log.Println(reflect.TypeOf(i))
		// for _, book := range books {
		// 	if book.ID == i {
		// 		json.NewEncoder(w).Encode(&book)
		// 	}
		// }
		// Info (HTTP) | log.Println(params)
		// ======== End

		var book models.Book

		params := mux.Vars(r)
		i, _ := strconv.Atoi(params["id"])

		// Info |
		log.Println(reflect.TypeOf(i))

		rows := db.QueryRow("SELECT * FROM books WHERE id = $1", i)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		json.NewEncoder(w).Encode(books)

		// Status
		log.Println("Get a book")
	}
}

// -------------------------------------

// AddBook ...
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // This func works

		// ======== This code works with Postman
		// var book Book
		// _ = json.NewDecoder(r.Body).Decode(&book)
		// books = append(books, book)
		// json.NewEncoder(w).Encode(books)
		// ======== End

		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
		logFatal(err)

		json.NewEncoder(w).Encode(bookID)

		// Info | log.Println(book)

		// Status
		log.Println("New book is added")
	}
}

// -------------------------------------

// UpdateBook ...
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// ======== This code works with Postman
		// var book Book
		// _ = json.NewDecoder(r.Body).Decode(&book)
		// for i, item := range books {
		// 	if item.ID == book.ID {
		// 		books[i] = book
		// 	}
		// }
		// json.NewEncoder(w).Encode(books)
		// ======== End

		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)

		res, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

		rowsUpdated, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)

		// Status
		log.Println("The book is updated")
	}
}

// -------------------------------------

// RemoveBook ...
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Works

		// ======== This code works with Postman
		// params := mux.Vars(r)
		// id, _ := strconv.Atoi(params["ID"])

		// for i, item := range books {
		// 	if item.ID == id {
		// 		books = append(books[:i], books[i+1:]...)
		// 	}
		// }
		// json.NewEncoder(w).Encode(books)
		// ======== End

		params := mux.Vars(r)
		i, _ := strconv.Atoi(params["id"])

		res, err := db.Exec("DELETE FROM books WHERE id = $1", i)
		logFatal(err)

		rowsDeleted, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)

		// Status
		log.Println("The book is deleted")
	}
}
