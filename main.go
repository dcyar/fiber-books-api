package main

import (
	"github.com/dcyar/fiber-books-api/config"
	"github.com/dcyar/fiber-books-api/database"
	router "github.com/dcyar/fiber-books-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	database.ConnectDb()

	app := fiber.New()

	app.Static("/uploads", "./uploads")

	router.SetUpRoutes(app)
	router.SetUpApiRoutes(app)

	app.Use(cors.New())

	// 404 "Not Found"
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(config.Config("PORT")))
}
