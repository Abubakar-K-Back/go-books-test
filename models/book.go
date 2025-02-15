package models

import "gorm.io/gorm"

// Book represents a book entity
// @Description Book model
type Book struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required"`
}
