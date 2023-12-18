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
	price    float64 `josn:"price"`
}

var books = []book{
	{ID: "1", Title: "title1", Author: "author1", Quantity: 3, price: 33.4},
	{ID: "2", Title: "title2", Author: "author2", Quantity: 5, price: 43.4},
	{ID: "3", Title: "title3", Author: "author3", Quantity: 7, price: 53.4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {

}
