package database

import (
	"github.com/dcyar/fiber-books-api/config"
	"github.com/dcyar/fiber-books-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBConn *gorm.DB

func ConnectDb() {
	dsn := config.Config("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Database connected")
	db.AutoMigrate(&models.Book{}, &models.User{})

	DBConn = db
}
