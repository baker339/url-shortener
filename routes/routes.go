package routes

import (
	"url-shortener/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", handlers.ShortenURL)
	app.Get("/:shortCode", handlers.RedirectURL)
}
