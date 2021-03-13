// controllers/books.go

package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/segej87/bookstore-test/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
  var books []models.Book
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
}

