package controllers

import (
	"github.com/dcyar/fiber-books-api/database"
	"github.com/dcyar/fiber-books-api/models"
	"github.com/dcyar/fiber-books-api/utils"
	"github.com/gofiber/fiber/v2"
)

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error on register request.",
		})
	}

	if err := database.DBConn.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists.",
			"data":    err,
		})
	}

	user.Password, _ = utils.HashPassword(user.Password)
	if err := database.DBConn.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't create a user.",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User was created successfully.",
	})
}

func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error on login request.",
			"data":    err,
		})
	}

	user := new(models.User)
	if err := database.DBConn.Where("email = ?", form.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User does not exists.",
			"data":    err,
		})
	}

	if !utils.CheckPasswordHash(form.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email or password are wrong.",
		})
	}

	token, err := utils.GenerateJwtToken(form.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
		"data": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
