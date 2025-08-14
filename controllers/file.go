package controllers

import (
	"redesign_backend/database"
	"redesign_backend/models"
	"redesign_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateFile(c *fiber.Ctx) error {
	var file models.File
	fileName := c.FormValue("file_name")
	filePath, err := utils.SaveFile(c, "file", false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if fileName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file_name is required",
		})
	}

	file.FileName = fileName
	file.FilePath = filePath

	if err := database.DB.Create(&file).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "File created successfully",
		"file":  file,
	})
}

func GetFile(c *fiber.Ctx) error {
	var file []models.File

	if err := database.DB.Find(&file).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get File",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"file": file,
	})
}