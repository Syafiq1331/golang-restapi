package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
var books []models.Book
models.DB.Find(&books)

c.JSON(http.StatusOK, gin.H{"data": books})
}
