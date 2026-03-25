package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var Authors []models.Author
var AuthorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
 
	result := []models.Author{}
	result = append(result, Authors...)
	json.NewEncoder(w).Encode(Authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if author.Name == "" {
		http.Error(w, "Invalid data", 400)
		return
	}

	author.ID = AuthorID
	AuthorID++
	Authors = append(Authors, author)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
	
}