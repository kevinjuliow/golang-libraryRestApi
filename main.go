package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevinjuliow/golang-libraryRestApi/config"
)

func main() {
	router := gin.Default()
	config.LoadEnv()
	dbPort := os.Getenv("PORT")

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Library rest APi")
	})

	router.Run("localhost:" + dbPort)
}
