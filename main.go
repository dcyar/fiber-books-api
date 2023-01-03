package main

import (
	"github.com/dcyar/fiber-books-api/database"
	router "github.com/dcyar/fiber-books-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"net/http"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	router.SetUpRoutes(app)
	router.SetUpApiRoutes(app)

	app.Use(cors.New())

	// 404 "Not Found"
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
