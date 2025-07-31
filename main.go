package main

import (
	"Backend/config"
	"Backend/router"
	"fmt"
	"log"
	"os"
	"strings"

	_ "Backend/docs" // Wajib agar swagger.json digunakan

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

// @title Backend Artikel API
// @version 1.0
// @description Dokumentasi REST API untuk manajemen artikel, kategori, komentar, dan penulis.
// @termsOfService http://swagger.io/terms/
// @contact.name Developer API Support
// @contact.email kamu@email.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @schemes http https
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Masukkan token JWT dengan format: Bearer {token}
func main() {
	config.DB = config.MongoConnect(config.DBName)
	if config.DB == nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	app := fiber.New()

	app.Use(logger.New())

	// Enhanced CORS configuration for both HTTP and HTTPS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, Access-Control-Allow-Origin",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin",
	}))

	// Add preflight OPTIONS handler
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(204)
	})

	app.Get("/docs/*", swagger.HandlerDefault) // http://localhost:6969/docs/index.html

	// Serve swagger.json with proper CORS at root level
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		return c.SendFile("./docs/swagger.json")
	})

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	log.Fatal(app.Listen(":" + port))
}
