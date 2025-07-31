package main

import (
	"Backend/config"
	"Backend/router"
	"fmt"
	"log"
	"os"

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
<<<<<<< HEAD
// @BasePath /
// @securityDefinitions.apikey Bearer
=======
// @securityDefinitions.apikey BearerAuth
>>>>>>> 36605f5109d3743dcee478d3c815d3b15f6f91d5
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

<<<<<<< HEAD
	app.Use(logger.New())

	// Enhanced CORS configuration for both HTTP and HTTPS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, Access-Control-Allow-Origin",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin",
=======
	// CORS harus paling atas
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://petstore.swagger.io,https://editor.swagger.io,https://app.swaggerhub.com,https://backendblog.up.railway.app,https://tampilan-blog.vercel.app",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
		ExposeHeaders: "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Authorization",
>>>>>>> 36605f5109d3743dcee478d3c815d3b15f6f91d5
	}))
	app.Use(logger.New())

<<<<<<< HEAD
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
=======
	// Handler global untuk preflight OPTIONS
	app.Options("*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", c.Get("Origin"))
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Get("/docs/*", swagger.HandlerDefault) // http://localhost:6969/docs/index.html
>>>>>>> 36605f5109d3743dcee478d3c815d3b15f6f91d5

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
