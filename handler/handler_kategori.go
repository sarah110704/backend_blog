package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllKategoris godoc
// @Summary Mendapatkan semua kategori
// @Description Mengambil semua data kategori dari database (butuh token)
// @Tags Kategori
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kategoris [get]
func GetAllKategoris(c *fiber.Ctx) error {
	kategoris, err := controller.GetAllKategoris(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data kategori",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua kategori",
		"data":    kategoris,
	})
}

// CreateKategori godoc
// @Summary Membuat kategori baru
// @Description Menambahkan kategori baru ke database (butuh token)
// @Tags Kategori
// @Accept json
// @Produce json
// @Param kategori body model.Kategori true "Data kategori baru"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/kategoris [post]
func CreateKategori(c *fiber.Ctx) error {
	var kategori model.Kategori

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	kategori.ID = primitive.NewObjectID().Hex()

	if err := controller.CreateKategori(&kategori); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan kategori",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kategori berhasil dibuat",
		"data":    kategori,
	})
}

// GetKategoriByID godoc
// @Summary Mendapatkan kategori berdasarkan ID
// @Description Mengambil detail kategori berdasarkan ID (butuh token)
// @Tags Kategori
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ID Kategori"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/kategoris/{id} [get]
func GetKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")

	kategori, err := controller.GetKategoriByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Kategori ditemukan",
		"data":    kategori,
	})
}

// UpdateKategoriByID godoc
// @Summary Mengupdate kategori berdasarkan ID
// @Description Mengubah data kategori berdasarkan ID (butuh token)
// @Tags Kategori
// @Accept json
// @Produce json
// @Param id path string true "ID Kategori"
// @Param kategori body model.Kategori true "Data kategori yang diupdate"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/kategoris/{id} [put]
func UpdateKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Update ID:", id)

	var kategori model.Kategori

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	err := controller.UpdateKategoriByID(c.Context(), id, kategori)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update kategori: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Kategori berhasil diperbarui",
	})
}

// DeleteKategoriByID godoc
// @Summary Menghapus kategori berdasarkan ID
// @Description Menghapus data kategori dari database berdasarkan ID (butuh token)
// @Tags Kategori
// @Accept json
// @Produce json
// @Param id path string true "ID Kategori"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/kategoris/{id} [delete]
func DeleteKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Hapus kategori ID:", id)

	err := controller.DeleteKategoriByID(c.Context(), id)
	if err != nil {
		fmt.Println("Delete kategori error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus kategori: %v", err),
		})
	}

	fmt.Println("Kategori berhasil dihapus")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Kategori dengan ID %s berhasil dihapus", id),
	})
}
