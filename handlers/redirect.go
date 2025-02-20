package handlers

import (
	"context"
	"database/sql"
	"url-shortener/config"

	"github.com/gofiber/fiber/v2"
)

func RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortCode")

	// Check Redis first
	url, err := config.RedisClient.Get(context.Background(), shortCode).Result()
	if err == nil {
		return c.Redirect(url, 301)
	}

	// Fetch from database if not in cache
	var originalURL string
	err = config.DB.QueryRow("SELECT original_url FROM urls WHERE short_code = $1", shortCode).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// Cache result in Redis
	config.RedisClient.Set(context.Background(), shortCode, originalURL, 0)

	// Redirect
	return c.Redirect(originalURL, 301)
}
