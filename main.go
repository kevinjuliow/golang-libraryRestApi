package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevinjuliow/golang-libraryRestApi/config"
	"github.com/kevinjuliow/golang-libraryRestApi/controllers"
)

func main() {
	router := gin.Default()
	config.LoadEnv()
	config.Dbconnection()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Library rest APi")
	})

	router.GET("/api/books", controllers.GetAllBooks)
	router.GET("/api/books/:id", controllers.GetById)
	router.POST("/api/books/", controllers.Create)
	router.PUT("/api/books/:id", controllers.Update)
	router.DELETE("/api/books/:id", controllers.Delete)

	//Port
	dbPort := os.Getenv("PORT")
	router.Run("localhost:" + dbPort)
}
