package main

import (
    "github.com/gin-gonic/gin"
    "restapi/handlers"
)

func main() {
    r := gin.Default()

    r.GET("/users", handlers.GetUsers)
    r.POST("/users", handlers.CreateUser)

    r.Run(":8080")
}