// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {  // Get model if exist
  var book models.Book

  if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": book})
}