package routes

import (
	"github.com/abkawan/go-books-api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up routes for the API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// @Summary Get all books
	// @Description Get a paginated list of books
	// @Accept  json
	// @Produce  json
	// @Param   limit  query  int  false  "Number of books to retrieve"
	// @Param   offset query  int  false  "Offset for pagination"
	// @Success 200 {array} models.Book
	// @Router /books [get]
	r.GET("/books", controllers.GetBooks)

	// @Summary Get a book by ID
	// @Description Retrieve details of a specific book by ID
	// @Accept  json
	// @Produce  json
	// @Param   id  path  int  true  "Book ID"
	// @Success 200 {object} models.Book
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [get]
	r.GET("/books/:id", controllers.GetBookByID)

	// @Summary Create a new book
	// @Description Add a new book to the database
	// @Accept  json
	// @Produce  json
	// @Param   book  body  models.Book  true  "Book data"
	// @Success 201 {object} models.Book
	// @Failure 400 {object} gin.H "Invalid request data"
	// @Router /books [post]
	r.POST("/books", controllers.CreateBook)

	// @Summary Update a book
	// @Description Update an existing book's details
	// @Accept  json
	// @Produce  json
	// @Param   id  path  int  true  "Book ID"
	// @Param   book  body  models.Book  true  "Updated book data"
	// @Success 200 {object} models.Book
	// @Failure 400 {object} gin.H "Invalid request data"
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [put]
	r.PUT("/books/:id", controllers.UpdateBook)

	// @Summary Delete a book
	// @Description Delete a book by ID
	// @Param   id  path  int  true  "Book ID"
	// @Success 200 {object} gin.H "Book deleted"
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [delete]
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}
