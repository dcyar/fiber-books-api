package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpApiRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	authRouter(api)
	authorsRouter(api)
	booksRouter(api)
}
