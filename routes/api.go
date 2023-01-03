package routes

import (
	c "github.com/dcyar/fiber-books-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	books := api.Group("/books")
	books.Get("/", c.BookList)
	books.Post("/", c.BookStore)
	books.Get("/:id", c.BookFind)
	books.Put("/:id", c.BookUpdate)
	books.Delete("/:id", c.BookDelete)
}
