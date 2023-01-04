package routes

import (
	c "github.com/dcyar/fiber-books-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func authRouter(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/register", c.Register)
	auth.Post("/login", c.Login)
}
