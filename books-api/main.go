package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Quantity int     `json:"Quantity"`
	Price    float64 `josn:"price"`
}

var books = []book{
	{ID: "1", Title: "title1", Author: "author1", Quantity: 3, Price: 33.4},
	{ID: "2", Title: "title2", Author: "author2", Quantity: 5, Price: 43.4},
	{ID: "3", Title: "title3", Author: "author3", Quantity: 7, Price: 53.4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
}
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookByID)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
