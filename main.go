package main

import (
	"go-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)
	r.POST("/books", handlers.CreateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.Run(":8080")
}
