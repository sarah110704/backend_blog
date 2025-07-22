package router

import (
	"Backend/handler"

	_ "Backend/docs" // penting agar swagger baca docs.go

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	// Route dokumentasi Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // akses di: http://localhost:5000/docs/index.html

	api := app.Group("/api")

	// Routing untuk homepage
	api.Get("/", handler.Homepage)

	// Routing untuk artikel
	api.Get("/artikels", handler.GetAllArtikels)
	api.Post("/artikels", handler.CreateArtikel)
	api.Get("/artikels/:id", handler.GetArtikelByID)
	api.Put("/artikels/:id", handler.UpdateArtikelByID)
	api.Delete("/artikels/:id", handler.DeleteArtikelByID)

	// Routing untuk kategori
	api.Get("/kategoris", handler.GetAllKategoris)
	api.Post("/kategoris", handler.CreateKategori)
	api.Get("/kategoris/:id", handler.GetKategoriByID)
	api.Put("/kategoris/:id", handler.UpdateKategoriByID)
	api.Delete("/kategoris/:id", handler.DeleteKategoriByID)

	// Routing untuk komentar
	api.Get("/komentars", handler.GetAllKomentars)
	api.Post("/komentars", handler.CreateKomentar)
	api.Get("/komentars/:id", handler.GetKomentarByID)
	api.Put("/komentars/:id", handler.UpdateKomentarByID)
	api.Delete("/komentars/:id", handler.DeleteKomentarByID)

	// Routing untuk penulis
	api.Get("/penulis", handler.GetAllPenulis)
	api.Post("/penulis", handler.CreatePenulis)
	api.Get("/penulis/:id", handler.GetPenulisByID)
	api.Put("/penulis/:id", handler.UpdatePenulisByID)
	api.Delete("/penulis/:id", handler.DeletePenulisByID)

	// Routing untuk auth
	// RegisterUser godoc
	// @Summary Register user
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User data to register"
	// @Success 201 {object} models.User
	// @Failure 400 {object} fiber.Map
	// @Router /api/register [post]
	api.Post("/register", handler.RegisterUser)
	api.Post("/login", handler.LoginUser)

}
