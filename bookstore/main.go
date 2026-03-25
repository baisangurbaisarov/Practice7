package main

import (
	"net/http"

	"bookstore/handlers"
)

func main() {

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetBooks(w, r)
		} else if r.Method == "POST" {
			handlers.CreateBook(w, r)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetBookByID(w, r)
		} else if r.Method == "PUT" {
			handlers.UpdateBook(w, r)
		} else if r.Method == "DELETE" {
			handlers.DeleteBook(w, r)
		}
	})

	http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetAuthors(w, r)
		} else if r.Method == "POST" {
			handlers.CreateAuthor(w, r)
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetCategories(w, r)
		} else if r.Method == "POST" {
			handlers.CreateCategory(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}