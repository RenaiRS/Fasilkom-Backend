package controllers

import (
	"redesign_backend/cache"
    "redesign_backend/database"
    "redesign_backend/models"
    "redesign_backend/utils"
    "time"
    "strconv"

    "github.com/gofiber/fiber/v2"
)

// controllers/scholarship.go
func GetScholarship(c *fiber.Ctx) error {
    limit := c.Query("limit")
    cacheKey := "scholarship_all"

    if cached, found := cache.GetScholarshipCache(cacheKey); found {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "data":  cached,
            "cache": true,
        })
    }

    var data []models.Scholarship
    query := database.DB.Model(&models.Scholarship{}).Order("tanggal DESC")
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

    cache.SetScholarshipCache(cacheKey, data)

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data":  data,
        "cache": false,
    })
}

func CreateScholarship(c *fiber.Ctx) error {
	var data models.Scholarship

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

func UpdateScholarship(c *fiber.Ctx) error {
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

	// Cek apakah scholarship ada
	var scholarship models.Scholarship
	if err := database.DB.First(&scholarship, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Scholarship tidak ditemukan",
		})
	}

	// Ambil data dari body
	var updateData models.Scholarship
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update data
	if err := database.DB.Model(&scholarship).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update scholarship",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Scholarship berhasil diperbarui",
		"data":    scholarship,
	})
}


func DeleteScholarship(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID scholarship is required",
		})
	}

	var item models.Scholarship
	if err := database.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Scholarship not found",
		})
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete scholarship",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
