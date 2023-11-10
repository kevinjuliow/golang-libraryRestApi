package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinjuliow/golang-libraryRestApi/config"
	"github.com/kevinjuliow/golang-libraryRestApi/models"
	"gorm.io/gorm"
)

type requestBody struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"Books": books})
}

func GetById(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message:": "Book Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error})
			return
		}
	}

	c.JSON(http.StatusOK, book)
}

func Create(c *gin.Context) {
	body := requestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Where("id = ?", body.ID).First(&models.Book{}).Error; err != gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"Error": "ID already Exists"})
		return
	}

	var book models.Book

	book.ID = body.ID
	book.Author = body.Author
	book.Title = body.Title
	book.Year = body.Year

	if err := config.DB.Create(&book).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Books": book})

}

func Update(c *gin.Context) {
	var book models.Book
	body := requestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "Id not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
	}

	if body.Author != "" {
		book.Author = body.Author
	}

	if body.Title != "" {
		book.Title = body.Title
	}

	if body.Year != 0 {
		book.Year = body.Year
	}

	config.DB.Save(&book)

}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "ID Not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
	}

	config.DB.Delete(book, id)
}
