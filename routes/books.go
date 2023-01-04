package routes

import (
	c "github.com/dcyar/fiber-books-api/controllers"
	"github.com/dcyar/fiber-books-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func booksRouter(api fiber.Router) {
	books := api.Group("/books")
	books.Get("/", c.BookList)
	books.Post("/", middleware.Protected(), c.BookStore)
	books.Get("/:id", c.BookFind)
	books.Put("/:id", middleware.Protected(), c.BookUpdate)
	books.Delete("/:id", middleware.Protected(), c.BookDelete)
}
