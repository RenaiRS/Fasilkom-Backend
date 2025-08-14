package controllers

import (
    "redesign_backend/database"
    "redesign_backend/models"
    "redesign_backend/utils"
    "time"
    "strconv"

    "github.com/gofiber/fiber/v2"
)

func GetVacancies(c *fiber.Ctx) error {
    var data []models.Vacancies
    limit := c.Query("limit")

    query := database.DB.Model(&models.Vacancies{}).Order("tanggal DESC")
    if limit != "" {
        if l, err := strconv.Atoi(limit); err == nil {
            query = query.Limit(l)
        }
    }

    if err := query.Find(&data).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal ambil data",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data": data,
    })
}

func CreateVacancies(c *fiber.Ctx) error {
	var data models.Vacancies

	// Ambil admin dari cookie token
	cookieToken := c.Cookies("admin_token")
	var admin models.Admin
	if err := database.DB.Where("token = ? AND token_exp >= ?", cookieToken, time.Now()).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}

	judul := c.FormValue("judul")
	if judul == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Judul is required",
		})
	}

	cover, err := utils.SaveFile(c, "cover", true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	postedAtStr := c.FormValue("posted_at")
	if postedAtStr != "" {
		t, err := time.Parse(time.RFC3339, postedAtStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid posted_at format. Use ISO 8601 (RFC3339)",
			})
		}
		data.Tanggal = t
	} else {
		data.Tanggal = time.Now()
	}

	data.Judul = judul
	data.Cover = cover
	data.Content = c.FormValue("content")
	data.PostedBy = admin.Username

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data created successfully",
		"data":    data,
	})
}

func UpdateVacancies(c *fiber.Ctx) error {
	// Cek token admin dari cookie
	cookieToken := c.Cookies("admin_token")
	var admin models.Admin
	if err := database.DB.Where("token = ? AND token_exp >= ?", cookieToken, time.Now()).
		First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}

	id := c.Params("id")

	// Cek apakah vacancies ada
	var vacancies models.Vacancies
	if err := database.DB.First(&vacancies, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vacancies tidak ditemukan",
		})
	}

	// Ambil data dari body
	var updateData models.Vacancies
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update data
	if err := database.DB.Model(&vacancies).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update vacancies",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Vacancies berhasil diperbarui",
		"data":    vacancies,
	})
}

func DeleteVacancies(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID vacancies is required",
		})
	}

	var item models.Vacancies
	if err := database.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vacancies not found",
		})
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete vacancies",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}