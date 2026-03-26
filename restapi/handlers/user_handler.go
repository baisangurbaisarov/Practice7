package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "restapi/models"
)

var users = []models.User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// GET: Retrieve all users
func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

// POST: Create a new user
func CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}