package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
	"log"
	"encoding/json"
)

type Book struct {
	ID		string `json:"id"`
	ISBN	string `json:"isbm"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}

type Author struct {
	FirstName	string `json:"firstname"`
	LastName 	string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, b := range books{
		if b.ID == params["id"] {
			json.NewEncoder(w).Encode(b)
			return
		} else {
			json.NewEncoder(w).Encode("nothing matched")
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode("success")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, b := range books {
		if b.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, b := range books {
		if b.ID == params["id"] {
			var book Book
			books = append(books[:index], books[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID: "1",
		ISBN: "p38thgw8g9hhdfl;idhsf;8",
		Title: "Book 1",
		Author: &Author{
			FirstName: "John:",
			LastName: "Pape",

		},
	})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

