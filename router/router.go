package router

import (
	"Backend/handler"
	"Backend/middleware"

	_ "Backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// Swagger Docs (http://localhost:6969/docs/index.html)
	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")

	// 🔓 Public Routes
	api.Get("/", handler.Homepage)
	api.Post("/login", handler.LoginUser)
	api.Post("/register", handler.RegisterUser)

	// Public GET for all resources
	api.Get("/artikels", handler.GetAllArtikels)
	api.Get("/artikels/:id", handler.GetArtikelByID)

	api.Get("/kategoris", handler.GetAllKategoris)
	api.Get("/kategoris/:id", handler.GetKategoriByID)

	api.Get("/komentars", handler.GetAllKomentars)
	api.Get("/komentars/:id", handler.GetKomentarByID)

	// 🔐 Protected Routes
	auth := api.Group("/", middleware.JWTProtected())

	// Artikel (POST, PUT, DELETE)
	auth.Post("/artikels", handler.CreateArtikel)
	auth.Put("/artikels/:id", handler.UpdateArtikelByID)
	auth.Delete("/artikels/:id", handler.DeleteArtikelByID)

	// Kategori
	auth.Post("/kategoris", handler.CreateKategori)
	auth.Put("/kategoris/:id", handler.UpdateKategoriByID)
	auth.Delete("/kategoris/:id", handler.DeleteKategoriByID)

	// Komentar
	auth.Post("/komentars", handler.CreateKomentar)
	auth.Put("/komentars/:id", handler.UpdateKomentarByID)
	auth.Delete("/komentars/:id", handler.DeleteKomentarByID)

	// Penulis (semua method harus pakai token)
	auth.Get("/penulis", handler.GetAllPenulis)
	auth.Get("/penulis/:id", handler.GetPenulisByID)
	auth.Post("/penulis", handler.CreatePenulis)
	auth.Put("/penulis/:id", handler.UpdatePenulisByID)
	auth.Delete("/penulis/:id", handler.DeletePenulisByID)
}
