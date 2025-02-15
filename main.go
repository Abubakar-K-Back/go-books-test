package main

import (
	"log"
	"os"

	_ "github.com/abkawan/go-books-api/docs"

	"github.com/abkawan/go-books-api/database"
	"github.com/abkawan/go-books-api/routes"

	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files" // ✅ FIXED IMPORT
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book Management API
// @version 1.0
// @description A REST API for managing books using Gin, PostgreSQL, Kafka, and Redis.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	r := routes.SetupRouter()

	// ✅ Add Swagger Route (Now Fixed)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
