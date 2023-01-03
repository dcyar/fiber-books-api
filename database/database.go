package database

import (
	"github.com/dcyar/fiber-books-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBConn *gorm.DB

func ConnectDb() {
	dsn := "root:password@tcp(127.0.0.1:3306)/fiber_books?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Database connected")
	db.AutoMigrate(&models.Book{})

	DBConn = db
}
