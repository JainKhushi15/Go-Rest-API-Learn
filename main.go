package main

import (
	//Standard Library to provide HTTP Client and Server Implementations
	"net/http"

	//Import path for the GIN Web Framework - Provides higher level API for building web applications in Go
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/books", lstBooksHndler)
	r.POST("/books", crtBookHandler)
	r.DELETE("/books/:id", delBookHandler)

	r.Run()
}

// Represents the structure of a book with 3 Fields
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// books variable is a slice of Book struct that holds the collection of books
var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

/*
Handles the GET Request
Responds with a status code of 200(OK)
*/
func lstBooksHndler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

/*
Handles the POST Request for creating a new book
Binds the JSON Payload in the request body to a Book struct
using ShouldBindJSON; Adds book to the Books slice
Responds with the created book as a JSON response - status code of 201(Created)
*/
func crtBookHandler(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

/*
Handles DELETE Request for deleting a Book based on the id Parameter
Iterates and removes the book with a matching ID
Responds with a status code of 204(No Content) to indicate successful deletion
*/
func delBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	ctx.Status(http.StatusNoContent)
}
