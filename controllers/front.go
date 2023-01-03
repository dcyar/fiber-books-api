package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func HomePage(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(Response{
		Status:  true,
		Message: "Welcome to Fiber Books API",
	})
}
