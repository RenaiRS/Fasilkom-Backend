package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string, isCover bool) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	if !isCover {
		allowedExt[".pdf"] = true
	}

	if !allowedExt[ext] {
		return "", fmt.Errorf("invalid file type")
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := "./assets/files/"
	publicPath := "/assets/files/"
	if isCover {
		saveDir = "./assets/covers/"
		publicPath = "/assets/covers/"
	}

	savePath := filepath.Join(saveDir, newFilename)
	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	fullURL := fmt.Sprintf("%s%s%s", c.BaseURL(), publicPath, newFilename)

	return fullURL, nil
}