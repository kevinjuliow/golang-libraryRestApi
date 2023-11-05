package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinjuliow/golang-libraryRestApi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env files")
	}
}

func Dbconnection() {
	var Book models.Book

	LoadEnv()

	dbPass := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")

	dbUrl := dbUsername + ":" + dbPass + "@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database : " + err.Error())
	}

	db.AutoMigrate(&Book)
}
