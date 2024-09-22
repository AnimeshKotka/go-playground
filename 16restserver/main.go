package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Book struct {
	BookId   string  `json:"bookId"`
	BookName string  `json:"bookName"`
	BooKdesc string  `json:"desc"`
	Price    int     `json:"price"`
	Author   *Author `json:author`
}

type Author struct {
	Name string `json:"name"`
}

// Fake DB
var books []Book

// Middleware/helpers
func (book *Book) IsEmpty() bool {
	return book.BookId == "" && book.BookName == ""
}

func main() {

	books = append(books, Book{BookId: "1", BookName: "Book 1", Price: 100, Author: &Author{Name: "Author 1"}})
	books = append(books, Book{BookId: "2", BookName: "Book 2", Price: 200, Author: &Author{Name: "Author 2"}})

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/books", getBooks).Methods("GET")

	r.HandleFunc("/book/{id}", getBookById).Methods("GET")

	r.HandleFunc("/book", addBook).Methods("Post")

	r.HandleFunc("/book/{id}", updateBook).Methods("Update")

	http.ListenAndServe(":3000", r)

}

//Controllers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Library</h1>"))
}
func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := r.Body

	if body == nil {
		panic("Request body is empty")
	}

	book := &Book{}
	err := json.NewDecoder(body).Decode(book)
	book.BookId = uuid.NewString()

	if err != nil {
		panic(err)
	}

	if book.IsEmpty() {
		panic("Empty book name")
	}

	books = append(books, *book)

	json.NewEncoder(w).Encode(book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, book := range books {
		if book.BookId == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Book not found"))
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, book := range books {
		if book.BookId == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.BookId = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Book not found"))
}
