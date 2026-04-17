package handlers

import (
    "net/http"
    "strconv"
    "time"

    "bookstoreGin/models"

    "github.com/gin-gonic/gin"
)

var Favorites []models.FavoriteBook

func GetFavoriteBooks(c *gin.Context) {
    userID := c.GetInt("user_id")

    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "5")
    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    if page <= 0 {
        page = 1
    }
    if limit <= 0 {
        limit = 5
    }

    favSet := map[int]bool{}
    for _, f := range Favorites {
        if f.UserID == userID {
            favSet[f.BookID] = true
        }
    }

    var favBooks []models.Book
    for _, b := range Books {
        if favSet[b.ID] {
            favBooks = append(favBooks, b)
        }
    }

    start := (page - 1) * limit
    end := start + limit
    if start > len(favBooks) {
        start = len(favBooks)
    }
    if end > len(favBooks) {
        end = len(favBooks)
    }

    result := favBooks[start:end]
    if result == nil {
        result = []models.Book{}
    }

    c.JSON(http.StatusOK, result)
}

func AddFavoriteBook(c *gin.Context) {
    userID := c.GetInt("user_id")

    bookID, err := strconv.Atoi(c.Param("bookId"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bookId"})
        return
    }

	bookFound := false
    for _, b := range Books {
        if b.ID == bookID {
            bookFound = true
            break
        }
    }
    if !bookFound {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    for _, f := range Favorites {
        if f.UserID == userID && f.BookID == bookID {
            c.JSON(http.StatusOK, gin.H{"message": "Already in favorites"})
            return
        }
    }

    entry := models.FavoriteBook{
        UserID:    userID,
        BookID:    bookID,
        CreatedAt: time.Now(),
    }
    Favorites = append(Favorites, entry)

    c.JSON(http.StatusOK, entry)
}

func RemoveFavoriteBook(c *gin.Context) {
    userID := c.GetInt("user_id")

    bookID, err := strconv.Atoi(c.Param("bookId"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bookId"})
        return
    }

    for i, f := range Favorites {
        if f.UserID == userID && f.BookID == bookID {
            Favorites = append(Favorites[:i], Favorites[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
}