package controllers

import (
	"errors"
	"github.com/dcyar/fiber-books-api/database"
	"github.com/dcyar/fiber-books-api/models"
	"github.com/dcyar/fiber-books-api/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func BookList(c *fiber.Ctx) error {
	books := []models.Book{}

	if err := database.DBConn.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books list",
		"data":    books,
	})
}

func BookFind(c *fiber.Ctx) error {
	book, err := findBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func BookStore(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	author, err := findAuthorById(strconv.Itoa(book.AuthorID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	book.Author = author

	uploadedFile, err := utils.UploadFile(c, "covers", "cover", "jpg,png")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	book.Cover = uploadedFile["path"]

	if err := database.DBConn.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book was created.",
		"data":    book,
	})
}

func BookUpdate(c *fiber.Ctx) error {
	bookForm := new(models.Book)

	if err := c.BodyParser(bookForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	book, err := findBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	author, err := findAuthorById(strconv.Itoa(book.AuthorID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if _, err := c.FormFile("cover"); err == nil {
		uploadedFile, err := utils.UploadFile(c, "covers", "cover", "jpg,png")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := utils.RemoveFile(book.Cover); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		book.Cover = uploadedFile["path"]
	}

	bookForm.Author = author
	bookForm.Cover = book.Cover

	if err := database.DBConn.Model(&book).Where("id = ?", book.ID).Updates(bookForm).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated",
		"data":    book,
	})
}

func BookDelete(c *fiber.Ctx) error {
	book, err := findBookById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.RemoveFile(book.Cover); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DBConn.Where("id = ?", book.ID).Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func findBookById(paramId string) (models.Book, error) {
	book := models.Book{}
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return book, errors.New("Book id are invalid")
	}

	if err := database.DBConn.Preload("Author").First(&book, id).Error; err != nil {
		return book, errors.New(err.Error())
	}

	return book, nil
}
