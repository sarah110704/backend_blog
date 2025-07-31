package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllKomentars godoc
// @Summary Ambil semua komentar
// @Description Mengambil semua data komentar dari database (butuh token)
// @Tags Komentar
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/komentars [get]
func GetAllKomentars(c *fiber.Ctx) error {
	komentars, err := controller.GetAllKomentars(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil komentar",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua komentar",
		"data":    komentars,
	})
}

// CreateKomentar godoc
// @Summary Tambah komentar baru
// @Description Membuat data komentar baru (membutuhkan token)
// @Tags Komentar
// @Accept json
// @Produce json
// @Param komentar body model.Komentar true "Data Komentar"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/komentars [post]
func CreateKomentar(c *fiber.Ctx) error {
	var komentar model.Komentar

	if err := c.BodyParser(&komentar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing request body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreateKomentar(&komentar); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan komentar",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Komentar berhasil dibuat",
		"data":    komentar,
	})
}

// GetKomentarByID godoc
// @Summary Ambil komentar berdasarkan ID
// @Description Mengambil satu data komentar berdasarkan ID (butuh token)
// @Tags Komentar
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ID Komentar"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/komentars/{id} [get]
func GetKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")

	komentar, err := controller.GetKomentarByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Komentar ditemukan",
		"data":    komentar,
	})
}

// UpdateKomentarByID godoc
// @Summary Update komentar
// @Description Memperbarui data komentar berdasarkan ID (membutuhkan token)
// @Tags Komentar
// @Accept json
// @Produce json
// @Param id path string true "ID Komentar"
// @Param komentar body model.Komentar true "Data Komentar"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/komentars/{id} [put]
func UpdateKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Raw body:", string(c.Body()))

	var updatedKomentar model.Komentar

	if err := c.BodyParser(&updatedKomentar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	err := controller.UpdateKomentarByID(c.Context(), id, updatedKomentar)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update komentar: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Komentar berhasil diperbarui",
	})
}

// DeleteKomentarByID godoc
// @Summary Hapus komentar
// @Description Menghapus komentar berdasarkan ID (membutuhkan token)
// @Tags Komentar
// @Accept json
// @Produce json
// @Param id path string true "ID Komentar"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/komentars/{id} [delete]
func DeleteKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Hapus komentar ID:", id)

	err := controller.DeleteKomentarByID(c.Context(), id)
	if err != nil {
		fmt.Println("Delete error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus komentar: %v", err),
		})
	}

	fmt.Println("Komentar berhasil dihapus")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Komentar dengan ID %s berhasil dihapus", id),
	})
}
