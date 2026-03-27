package handlers

import (
	"net/http"

	"bookstoreGin/models"

	"github.com/gin-gonic/gin"
)

var Authors []models.Author
var AuthorID = 1

func GetAuthors(c *gin.Context) {
	result := []models.Author{}
	result = append(result, Authors...)
	c.JSON(http.StatusOK, result)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if author.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	author.ID = AuthorID
	AuthorID++
	Authors = append(Authors, author)

	c.JSON(http.StatusCreated, author)
}