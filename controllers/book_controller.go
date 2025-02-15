package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abkawan/go-books-api/database"
	"github.com/abkawan/go-books-api/kafka"
	"github.com/abkawan/go-books-api/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// ✅ GET /books (Supports Pagination + Redis Caching)
func GetBooks(c *gin.Context) {
	var books []models.Book

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit or offset"})
		return
	}

	cacheKey := "books:limit=" + limitStr + ":offset=" + offsetStr
	cachedBooks, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(cachedBooks), &books)
		if len(books) > 0 {
			c.JSON(http.StatusOK, books)
			return
		}
	}

	result := database.DB.Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	booksJSON, _ := json.Marshal(books)
	redisClient.Set(ctx, cacheKey, booksJSON, 300)

	c.JSON(http.StatusOK, books)
}

// ✅ GET /books/:id (Redis Cache)
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	cachedBook, err := redisClient.Get(ctx, "book:"+id).Result()
	if err == nil {
		var book models.Book
		json.Unmarshal([]byte(cachedBook), &book)
		c.JSON(http.StatusOK, book)
		return
	}

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	bookJSON, _ := json.Marshal(book)
	redisClient.Set(ctx, "book:"+id, bookJSON, 0)

	c.JSON(http.StatusOK, book)
}

// ✅ POST /books (Creates a Book + Kafka Event)
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)

	// ✅ Publish Kafka Event
	bookJSON, _ := json.Marshal(book)
	kafka.ProduceMessage("book_events", string(bookJSON))

	redisClient.Del(ctx, "books")
	c.JSON(http.StatusCreated, book)
}

// ✅ PUT /books/:id (Updates a Book + Kafka Event)
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&book)

	// ✅ Publish Kafka Event
	bookJSON, _ := json.Marshal(book)
	kafka.ProduceMessage("book_events", string(bookJSON))

	redisClient.Del(ctx, "books")
	redisClient.Del(ctx, "book:"+id)
	c.JSON(http.StatusOK, book)
}

// ✅ DELETE /books/:id (Deletes a Book + Kafka Event)
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Book{}, id)

	// ✅ Publish Kafka Event
	kafka.ProduceMessage("book_events", `{"event":"book_deleted","id":`+id+`}`)

	redisClient.Del(ctx, "books")
	redisClient.Del(ctx, "book:"+id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
