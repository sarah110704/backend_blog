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
	// Load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Println("Error loading .env file")
		}
	} else {
		log.Println(".env file not found, using environment variables from Railway")
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
// @host localhost:6969
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @BasePath /
// @name Authorization
// @description Masukkan token JWT dengan format: Bearer {token}
func main() {
	config.DB = config.MongoConnect(config.DBName)
	if config.DB == nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/docs/*", swagger.HandlerDefault) // http://localhost:6969/docs/index.html

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

	host := os.Getenv("RAILWAY_STATIC_URL")
	if host == "" {
		host = "http://localhost"
	}

	fmt.Printf("🚀 Server running at %s:%s\n", host, port)
	log.Fatal(app.Listen(":" + port))
}
