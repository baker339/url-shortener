package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"url-shortener/config"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize database and Redis
	config.InitDB()
	config.InitRedis()

	// Start Fiber app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Run server
	port := ":3000"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(app.Listen(port))
}
