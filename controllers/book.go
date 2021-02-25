package controllers

import (
	"net/http"

	"github.com/JenYata/goEx/models"
	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

//POST /books
//Create new book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil { //check JSON format is valid
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

//GET /books/:id
//Find a book
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(&models.Book{
		Title:  input.Title,
		Author: input.Author,
	})

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&book).Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
