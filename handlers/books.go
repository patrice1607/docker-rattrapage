package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books = []Book{
		{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan"},
		{ID: 2, Title: "Clean Code", Author: "Robert Martin"},
	}
	mu     sync.Mutex
	nextID = 3
)

func GetBooks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	book.ID = nextID
	nextID++
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
