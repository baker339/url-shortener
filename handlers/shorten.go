package handlers

import (
	"context"
	"fmt"
	"url-shortener/config"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
)

func ShortenURL(c *fiber.Ctx) error {
	var request models.URL
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Generate a unique short code
	shortCode := utils.GenerateShortCode(6)

	// Insert into database
	_, err := config.DB.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortCode, request.OriginalURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save URL"})
	}

	// Cache in Redis
	config.RedisClient.Set(context.Background(), shortCode, request.OriginalURL, 0)

	// Return response
	return c.JSON(fiber.Map{"short_url": fmt.Sprintf("http://localhost:3000/%s", shortCode)})
}
