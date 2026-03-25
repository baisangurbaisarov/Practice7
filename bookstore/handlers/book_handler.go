package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bookstore/models"
)

var Books []models.Book
var BookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	categoryStr := r.URL.Query().Get("category")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	categoryID, _ := strconv.Atoi(categoryStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	var filtered []models.Book
	for _, b := range Books {
		if categoryStr != "" && b.CategoryID != categoryID {
			continue
		}
		filtered = append(filtered, b)
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	json.NewEncoder(w).Encode(filtered[start:end])
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, b := range Books {
		if b.ID == id {
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)

    if book.Title == "" {
        http.Error(w, "Title is required", http.StatusBadRequest)
        return
    }
    if book.Price < 0.01 {
        http.Error(w, "Price must be at least 0.01", http.StatusBadRequest)
        return
    }

    authorFound := false
    for _, a := range Authors {
        if a.ID == book.AuthorID {
            authorFound = true
            break
        }
    }
    if !authorFound {
        http.Error(w, "Invalid author_id", http.StatusBadRequest)
        return
    }

    categoryFound := false
    for _, c := range Categories {
        if c.ID == book.CategoryID {
            categoryFound = true
            break
        }
    }
    if !categoryFound {
        http.Error(w, "Invalid category_id", http.StatusBadRequest)
        return
    }

    book.ID = BookID
    BookID++
    Books = append(Books, book)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i := range Books {
		if Books[i].ID == id {
			var updated models.Book
			json.NewDecoder(r.Body).Decode(&updated)

			if updated.Title == "" {
                http.Error(w, "Title is required", http.StatusBadRequest)
                return
            }
            if updated.Price < 0.01 {
                http.Error(w, "Price must be at least 0.01", http.StatusBadRequest)
                return
            }

            updated.ID = id
            Books[i] = updated
            json.NewEncoder(w).Encode(Books[i])
            return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i := range Books {
		if Books[i].ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}