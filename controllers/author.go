package controllers

import (
	"errors"
	"github.com/dcyar/fiber-books-api/database"
	"github.com/dcyar/fiber-books-api/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AuthorList(c *fiber.Ctx) error {
	authors := []models.Author{}

	if err := database.DBConn.Find(&authors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authors list",
		"data":    authors,
	})
}

func AuthorFind(c *fiber.Ctx) error {
	author, err := findAuthorById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

func AuthorStore(c *fiber.Ctx) error {
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.DBConn.Create(&author).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Author was created.",
		"data":    author,
	})
}

func AuthorUpdate(c *fiber.Ctx) error {
	authorForm := new(models.Author)

	if err := c.BodyParser(authorForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	author, err := findAuthorById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DBConn.Model(&author).Where("id = ?", author.ID).Updates(authorForm).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Author updated",
		"data":    author,
	})
}

func AuthorDelete(c *fiber.Ctx) error {
	author, err := findAuthorById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DBConn.Select("Books").Where("id = ?", author.ID).Delete(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func findAuthorById(paramId string) (models.Author, error) {
	author := models.Author{}
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return author, errors.New("Author id are invalid")
	}

	if err := database.DBConn.First(&author, id).Error; err != nil {
		return author, errors.New(err.Error())
	}

	return author, nil
}
