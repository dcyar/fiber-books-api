package controllers

import (
	"github.com/dcyar/fiber-books-api/database"
	"github.com/dcyar/fiber-books-api/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func BookList(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.Find(&books)

	return c.Status(http.StatusOK).JSON(struct {
		Message string        `json:"message"`
		Data    []models.Book `json:"data"`
	}{
		Message: "Books list",
		Data:    books,
	})
}

func BookFind(c *fiber.Ctx) error {
	book := models.Book{}

	database.DBConn.First(&book, c.Params("id"))

	return c.Status(http.StatusOK).JSON(book)
}

func BookStore(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.DBConn.Create(&book)

	return c.Status(http.StatusCreated).JSON(struct {
		Message string       `json:"message"`
		Data    *models.Book `json:"data"`
	}{
		Message: "Book was created.",
		Data:    book,
	})
}

func BookUpdate(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Model(&models.Book{}).Where("id = ?", id).Updates(models.Book{Title: book.Title, Author: book.Author, Year: book.Year})

	return c.Status(http.StatusOK).JSON(struct {
		Message string `json:"message"`
	}{Message: "Updated"})
}

func BookDelete(c *fiber.Ctx) error {
	book := new(models.Book)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Where("id = ?", id).Delete(&book)

	return c.SendStatus(http.StatusNoContent)
}
