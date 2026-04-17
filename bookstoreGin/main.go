package main

import (
    "bookstoreGin/handlers"
    "bookstoreGin/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBookByID)

    r.GET("/authors", handlers.GetAuthors)
    r.POST("/authors", handlers.CreateAuthor)
    r.GET("/categories", handlers.GetCategories)
    r.POST("/categories", handlers.CreateCategory)

    auth := r.Group("/")
    auth.Use(middleware.AuthRequired())
    {
        auth.POST("/books", handlers.CreateBook)
        auth.PUT("/books/:id", handlers.UpdateBook)
        auth.DELETE("/books/:id", handlers.DeleteBook)

        auth.GET("/books/favorites", handlers.GetFavoriteBooks)
        auth.PUT("/books/:bookId/favorites", handlers.AddFavoriteBook)
        auth.DELETE("/books/:bookId/favorites", handlers.RemoveFavoriteBook)
    }

    r.Run(":8080")
}