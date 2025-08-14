package controllers

import (
	"redesign_backend/cache"
	"redesign_backend/database"
	"redesign_backend/models"
	"redesign_backend/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetBerita(c *fiber.Ctx) error {
	isPriority := c.Query("is_priority")
	limit := c.Query("limit")
	cacheKey := "berita_all"

	if isPriority == "true" {
		cacheKey = "berita_priority"
	}

	if cached, found := cache.GetBeritaCache(cacheKey); found {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"berita": cached,
			"cache":  true,
		})
	}

	var berita []models.Berita
	query := database.DB.Model(&models.Berita{}).Where("tanggal <= ?", time.Now()).Order("tanggal DESC")
	if isPriority == "true" {
		query = query.Where("is_priority = ?", true)
	}
	if limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(limitInt)
		}
	}

	if err := query.Find(&berita).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Berita not found",
		})
	}

	cache.SetBeritaCache(cacheKey, berita)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"berita": berita,
		"cache":  false,
	})
}

func CreateBerita(c *fiber.Ctx) error {
	var berita models.Berita

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

	isPriorityStr := c.FormValue("is_priority")
	isPriority, err := strconv.ParseBool(isPriorityStr)
	if err != nil {
		isPriority = false
	}

	postedAtStr := c.FormValue("posted_at")
	if postedAtStr != "" {
		t, err := time.Parse(time.RFC3339, postedAtStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid posted_at format. Use ISO 8601 (RFC3339)",
			})
		}
		berita.Tanggal = t
	} else {
		berita.Tanggal = time.Now() // otomatis pakai waktu saat ini
	}

	berita.Judul = judul
	berita.Cover = cover
	berita.IsPriority = isPriority
	berita.Content = c.FormValue("content")
	berita.PostedBy = admin.Username 

	if err := database.DB.Create(&berita).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save berita",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berita created successfully",
		"berita":  berita,
	})
}

func UpdateBerita(c *fiber.Ctx) error {
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

    // Cek apakah berita ada
    var berita models.Berita
    if err := database.DB.First(&berita, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Berita tidak ditemukan",
        })
    }

    // Ambil data dari body
    var updateData models.Berita
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Update data
    if err := database.DB.Model(&berita).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal update berita",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Berita berhasil diperbarui",
        "data":    berita,
    })
}

// DeleteBerita menghapus berita berdasarkan ID
func DeleteBerita(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID berita is required",
		})
	}
	var berita models.Berita

	if err := database.DB.Where("id = ?", id).First(&berita).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Berita not found",
		})
	}

	if err := database.DB.Delete(&berita).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete berita",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
