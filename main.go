package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevinjuliow/golang-libraryRestApi/config"
	"github.com/kevinjuliow/golang-libraryRestApi/routes"
)

func main() {
	router := gin.Default()
	config.LoadEnv()
	config.Dbconnection()
	routes.Routes()

	//Port
	dbPort := os.Getenv("PORT")
	router.Run("localhost:" + dbPort)
}
