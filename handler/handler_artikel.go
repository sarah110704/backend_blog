package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Homepage godoc
// @Summary Menampilkan halaman utama
// @Description Menampilkan teks sambutan dari API
// @Tags Homepage
// @Accept json
// @Produce json
// @Success 200 {string} string "Welcome to the Home Page!"
// @Router /api/ [get]
func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Home Page!")
}

// GetAllArtikels godoc
// @Summary Mendapatkan semua artikel
// @Description Mengambil semua data artikel dari database
// @Tags Artikel
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/artikels [get]
func GetAllArtikels(c *fiber.Ctx) error {
	artikels, err := controller.GetAllArtikels(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch artikels",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Fetched all artikels",
		"data":    artikels,
	})
}

// CreateArtikel godoc
// @Summary Membuat artikel baru
// @Description Menyimpan artikel baru ke database (butuh token)
// @Tags Artikel
// @Accept json
// @Produce json
// @Param artikel body model.Artikel true "Data artikel baru"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/artikels [post]
func CreateArtikel(c *fiber.Ctx) error {
	var artikel model.Artikel

	if err := c.BodyParser(&artikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing request body",
			"error":   err.Error(),
		})
	}

	// Ambil id_penulis dari JWT (email)
	user := c.Locals("user")
	if user != nil {
		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		artikel.IDPenulis = claims["email"].(string)
	}

	// Jika ingin id_kategori otomatis, bisa set default di sini
	// artikel.IDKategori = "default_kategori_id"

	if err := controller.CreateArtikel(&artikel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan artikel",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Artikel berhasil dibuat",
		"data":    artikel,
	})
}

// GetArtikelByID godoc
// @Summary Mendapatkan artikel berdasarkan ID
// @Description Mengambil artikel tertentu berdasarkan ID
// @Tags Artikel
// @Accept json
// @Produce json
// @Param id path string true "ID Artikel"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/artikels/{id} [get]
func GetArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	artikel, err := controller.GetArtikelByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Artikel ditemukan",
		"data":    artikel,
	})
}

// UpdateArtikelByID godoc
// @Summary Mengupdate artikel berdasarkan ID
// @Description Mengubah data artikel berdasarkan ID (butuh token)
// @Tags Artikel
// @Accept json
// @Produce json
// @Param id path string true "ID Artikel"
// @Param artikel body model.Artikel true "Data artikel yang diupdate"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/artikels/{id} [put]
func UpdateArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedArtikel model.Artikel
	if err := c.BodyParser(&updatedArtikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	err := controller.UpdateArtikelByID(c.Context(), id, updatedArtikel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update artikel: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Artikel berhasil diperbarui",
		"data":    updatedArtikel,
	})
}

// DeleteArtikelByID godoc
// @Summary Menghapus artikel berdasarkan ID
// @Description Menghapus artikel tertentu dari database (butuh token)
// @Tags Artikel
// @Accept json
// @Produce json
// @Param id path string true "ID Artikel"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/artikels/{id} [delete]
func DeleteArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	err := controller.DeleteArtikelByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus artikel: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Artikel dengan ID %s berhasil dihapus", id),
	})
}
