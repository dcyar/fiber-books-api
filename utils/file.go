package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(c *fiber.Ctx, folder string, input string, validExtensions string) (map[string]string, error) {
	file, err := c.FormFile(input)

	if err != nil {
		return map[string]string{}, err
	}

	extension := filepath.Ext(file.Filename)
	if err := validCoverExtension(extension, validExtensions); err != nil {
		return map[string]string{}, err
	}

	if _, err := os.ReadDir("uploads/" + folder); err != nil {
		if err := os.MkdirAll("uploads/"+folder, 0755); err != nil {
			return map[string]string{}, err
		}
	}

	filePath := fmt.Sprintf("uploads/%s/%s%s", folder, uuid.New(), extension)
	if err := c.SaveFile(file, "./"+filePath); err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"path": filePath,
		"url":  fmt.Sprintf("%s/%s", c.BaseURL(), filePath),
	}, nil
}

func RemoveFile(path string) error {
	if err := os.Remove("./" + path); err != nil {
		return err
	}

	return nil
}

func validCoverExtension(extension string, extensions string) error {
	for _, ext := range strings.Split(extensions, ",") {
		if extension == "."+ext {
			return nil
		}
	}

	return errors.New("invalid file extension")
}
