package router

import (
	_ "Backend/docs"
	"Backend/handler"
	"Backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")

	api.Get("/", handler.Homepage)
	api.Post("/login", handler.LoginUser)
	api.Post("/register", handler.RegisterUser)

	protected := api.Group("/", middleware.JWTProtected())

	protected.Get("/artikels", handler.GetAllArtikels)
	protected.Get("/artikels/:id", handler.GetArtikelByID)
	protected.Post("/artikels", handler.CreateArtikel)
	protected.Put("/artikels/:id", handler.UpdateArtikelByID)
	protected.Delete("/artikels/:id", handler.DeleteArtikelByID)

	protected.Get("/kategoris", handler.GetAllKategoris)
	protected.Get("/kategoris/:id", handler.GetKategoriByID)
	protected.Post("/kategoris", handler.CreateKategori)
	protected.Put("/kategoris/:id", handler.UpdateKategoriByID)
	protected.Delete("/kategoris/:id", handler.DeleteKategoriByID)

	protected.Get("/komentars", handler.GetAllKomentars)
	protected.Get("/komentars/:id", handler.GetKomentarByID)
	protected.Post("/komentars", handler.CreateKomentar)
	protected.Put("/komentars/:id", handler.UpdateKomentarByID)
	protected.Delete("/komentars/:id", handler.DeleteKomentarByID)

	protected.Get("/penulis", handler.GetAllPenulis)
	protected.Get("/penulis/:id", handler.GetPenulisByID)
	protected.Post("/penulis", handler.CreatePenulis)
	protected.Put("/penulis/:id", handler.UpdatePenulisByID)
	protected.Delete("/penulis/:id", handler.DeletePenulisByID)
}
