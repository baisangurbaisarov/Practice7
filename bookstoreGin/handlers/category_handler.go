package handlers

import (
	"net/http"

	"bookstore/models"

	"github.com/gin-gonic/gin"
)

var Categories []models.Category
var CategoryID = 1

func GetCategories(c *gin.Context) {
	result := []models.Category{}
	result = append(result, Categories...)
	c.JSON(http.StatusOK, result)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	category.ID = CategoryID
	CategoryID++
	Categories = append(Categories, category)

	c.JSON(http.StatusCreated, category)
}