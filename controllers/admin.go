package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"redesign_backend/cache"
	"redesign_backend/database"
	"redesign_backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var admin models.Admin
	username := c.FormValue("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username is required",
		})
	}
	password := c.FormValue("password")
	if password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password is required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	admin.Username = username
	admin.Password = string(hashedPassword)

	if err := database.DB.Create(&admin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create admin",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Admin created successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var admin models.Admin
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and Password is required",
		})
	}

	if err := database.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong Credential",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong Credential",
		})
	}

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	token := hex.EncodeToString(bytes)

	expTime := time.Now().Add(1 * time.Hour) // simpan ke variabel dulu
	admin.Token = token
	admin.TokenExp = &expTime               // pakai pointer

	if err := database.DB.Save(&admin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save admin token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "admin_token",
		Value:    token,
		Expires:  expTime,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successfully",
	})
}

func Logout(c *fiber.Ctx) error {
	var admin models.Admin
	cookieToken := c.Cookies("admin_token")
	if err := database.DB.Where("token =? AND token_exp >= ?", cookieToken, time.Now()).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token is invalid",
		})
	}

	admin.Token = ""
	admin.TokenExp = nil // kosongkan pointer

	if err := database.DB.Save(&admin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save admin token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "admin_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout Successfully",
	})
}

func SyncBeritaCache(c *fiber.Ctx) error {
	var all []models.Berita
	if err := database.DB.Model(&models.Berita{}).Where("posted_at <= ?", time.Now()).Find(&all).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal ambil semua berita",
		})
	}

	var priority []models.Berita
	if err := database.DB.Model(&models.Berita{}).Where("posted_at <= ? AND is_priority = ?", time.Now(), true).Find(&priority).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal ambil berita priority",
		})
	}

	cache.SetBeritaCache("berita_all", all)
	cache.SetBeritaCache("berita_priority", priority)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cache berita berhasil disinkronisasi ulang",
	})
}
