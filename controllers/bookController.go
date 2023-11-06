package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinjuliow/golang-libraryRestApi/config"
	"github.com/kevinjuliow/golang-libraryRestApi/models"
	"gorm.io/gorm"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"Books": books})
}

func GetById(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	if err := config.DB.First(book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Book Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, book)
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
