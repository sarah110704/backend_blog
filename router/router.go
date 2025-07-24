package router

import (
	_ "Backend/docs"
	"Backend/handler"
	"Backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// Dokumentasi Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // http://localhost:6969/docs/index.html

	api := app.Group("/api")

	// 🔓 PUBLIC ROUTES
	api.Get("/", handler.Homepage)
	api.Post("/login", handler.LoginUser)
	api.Post("/register", handler.RegisterUser)

	// 🔐 PROTECTED ROUTES (Harus pakai JWT)
	auth := api.Group("/", middleware.JWTProtected())

	// Artikel
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

	// Penulis
	auth.Post("/penulis", handler.CreatePenulis)
	auth.Put("/penulis/:id", handler.UpdatePenulisByID)
	auth.Delete("/penulis/:id", handler.DeletePenulisByID)

	// 🔓 OPSIONAL: GET masih publik (boleh dibuka ke semua)
	api.Get("/artikels", handler.GetAllArtikels)
	api.Get("/artikels/:id", handler.GetArtikelByID)

	api.Get("/kategoris", handler.GetAllKategoris)
	api.Get("/kategoris/:id", handler.GetKategoriByID)

	api.Get("/komentars", handler.GetAllKomentars)
	api.Get("/komentars/:id", handler.GetKomentarByID)

	api.Get("/penulis", handler.GetAllPenulis)
	api.Get("/penulis/:id", handler.GetPenulisByID)
}
