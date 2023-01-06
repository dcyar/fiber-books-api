package routes

import (
	c "github.com/dcyar/fiber-books-api/controllers"
	"github.com/dcyar/fiber-books-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func authorsRouter(api fiber.Router) {
	authors := api.Group("/authors")
	authors.Get("/", c.AuthorList)
	authors.Post("/", middleware.Protected(), c.AuthorStore)
	authors.Get("/:id", c.AuthorFind)
	authors.Put("/:id", middleware.Protected(), c.AuthorUpdate)
	authors.Delete("/:id", middleware.Protected(), c.AuthorDelete)
}
