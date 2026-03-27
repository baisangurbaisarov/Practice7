package handlers

import (
	"net/http"
	"strconv"

	"bookstoreGin/models"

	"github.com/gin-gonic/gin"
)

var Books []models.Book
var BookID = 1

// GetBooks godoc
// GET /books — returns a paginated, optionally filtered list of books
// Query params:
//
//	page     int  (default 1)
//	limit    int  (default 5)
//	category int  (filter by CategoryID)
func GetBooks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")
	categoryStr := c.Query("category")

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

	result := filtered[start:end]
	if result == nil {
		result = []models.Book{}
	}

	c.JSON(http.StatusOK, result)
}

// GetBookByID godoc
// GET /books/:id — returns a single book by ID
func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, b := range Books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// CreateBook godoc
// POST /books — creates a new book
// Required fields: title, author_id (must exist), category_id (must exist), price (>= 0.01)
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if book.Price < 0.01 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be at least 0.01"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author_id"})
		return
	}

	categoryFound := false
	for _, cat := range Categories {
		if cat.ID == book.CategoryID {
			categoryFound = true
			break
		}
	}
	if !categoryFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
		return
	}

	book.ID = BookID
	BookID++
	Books = append(Books, book)

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updated.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if updated.Price < 0.01 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be at least 0.01"})
		return
	}

	for i := range Books {
		if Books[i].ID == id {
			updated.ID = id
			Books[i] = updated
			c.JSON(http.StatusOK, Books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i := range Books {
		if Books[i].ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}