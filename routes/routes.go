package routes

import (
	"github.com/abkawan/go-books-api/controllers" // ✅ Import models
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// @Summary Get all books
	// @Description Get a paginated list of books
	// @Tags Books
	// @Accept json
	// @Produce json
	// @Param limit query int false "Number of books to retrieve"
	// @Param offset query int false "Offset for pagination"
	// @Success 200 {array} models.Book
	// @Router /books [get]
	r.GET("/books", controllers.GetBooks)

	// @Summary Get a book by ID
	// @Description Retrieve a specific book
	// @Tags Books
	// @Accept json
	// @Produce json
	// @Param id path int true "Book ID"
	// @Success 200 {object} models.Book
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [get]
	r.GET("/books/:id", controllers.GetBookByID)

	// @Summary Create a book
	// @Description Add a new book to the database
	// @Tags Books
	// @Accept json
	// @Produce json
	// @Param book body models.Book true "Book data"
	// @Success 201 {object} models.Book
	// @Failure 400 {object} gin.H "Invalid request data"
	// @Router /books [post]
	r.POST("/books", controllers.CreateBook)

	// @Summary Update a book
	// @Description Modify an existing book
	// @Tags Books
	// @Accept json
	// @Produce json
	// @Param id path int true "Book ID"
	// @Param book body models.Book true "Updated book data"
	// @Success 200 {object} models.Book
	// @Failure 400 {object} gin.H "Invalid request data"
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [put]
	r.PUT("/books/:id", controllers.UpdateBook)

	// @Summary Delete a book
	// @Description Remove a book from the database
	// @Tags Books
	// @Param id path int true "Book ID"
	// @Success 200 {object} gin.H "Book deleted"
	// @Failure 404 {object} gin.H "Book not found"
	// @Router /books/{id} [delete]
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}
