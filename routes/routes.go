package routes

import (
	c "github.com/dcyar/fiber-books-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", c.HomePage)
}
