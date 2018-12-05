package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

// Book ...
type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   int    `json:year`
}

// Array of books
var books []Book

// DataBase
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

	pqUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	sql.Open("postgres", pqUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	// Info |
	log.Println(pqUrl)

	router := mux.NewRouter()

	// ======== Temp-Data ========

	// books = append(books, Book{ID: 1, Title: "Les Misérables", Author: "Victor Hugo", Year: 1862},
	// 	Book{ID: 2, Title: "I Promessi sposi", Author: "Alessandro Manzoni", Year: 1847},
	// 	Book{ID: 3, Title: "Les Blancs et les Bleus", Author: "Alexandre Dumas", Year: 1867},
	// 	Book{ID: 4, Title: "A Moveable Feast", Author: "Ernest Hemingway", Year: 1964},
	// 	Book{ID: 5, Title: "Drei Kameraden", Author: "Erich Maria Remarque", Year: 1936})

	// ======== Http-Actions ========

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{ID}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/book/{ID}", removeBook).Methods("DELETE")

	// ======== Server? ========

	log.Fatal(http.ListenAndServe(":8080", router))
}

// ======== Book-Actions ========

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Temp Code | json.NewEncoder(w).Encode(books)

	// Status
	log.Println("Gets all books")
}

// -------------------------------------

func getBook(w http.ResponseWriter, r *http.Request) {

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

	// Status
	log.Println("Get a books")
}

// -------------------------------------

func addBook(w http.ResponseWriter, r *http.Request) {

	// ======== This code works with Postman
	// var book Book
	// _ = json.NewDecoder(r.Body).Decode(&book)
	// books = append(books, book)
	// json.NewEncoder(w).Encode(books)
	// ======== End

	// Status
	log.Println("New book is added")
}

// -------------------------------------

func updateBook(w http.ResponseWriter, r *http.Request) {

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

	// Status
	log.Println("The book is updated")
}

// -------------------------------------

func removeBook(w http.ResponseWriter, r *http.Request) {

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

	// Status
	log.Println("The book is deleted")
}
