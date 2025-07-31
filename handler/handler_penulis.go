package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllPenulis godoc
// @Summary Ambil semua penulis
// @Description Mengambil semua data penulis dari database (butuh token)
// @Tags Penulis
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/penulis [get]
func GetAllPenulis(c *fiber.Ctx) error {
	penulis, err := controller.GetAllPenulis(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data penulis",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua data penulis",
		"data":    penulis,
	})
}

// CreatePenulis godoc
// @Summary Tambah penulis baru
// @Description Menambahkan data penulis ke database (dengan token Bearer)
// @Tags Penulis
// @Accept json
// @Produce json
// @Param penulis body model.Penulis true "Data Penulis"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/penulis [post]
func CreatePenulis(c *fiber.Ctx) error {
	var p model.Penulis

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreatePenulis(&p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat penulis",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Penulis berhasil dibuat",
		"data":    p,
	})
}

// GetPenulisByID godoc
// @Summary Ambil penulis berdasarkan ID
// @Description Mengambil data penulis berdasarkan ID (butuh token)
// @Tags Penulis
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "ID Penulis"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/penulis/{id} [get]
func GetPenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := controller.GetPenulisByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Penulis ditemukan",
		"data":    p,
	})
}

// UpdatePenulisByID godoc
// @Summary Perbarui data penulis
// @Description Memperbarui data penulis berdasarkan ID (dengan token Bearer)
// @Tags Penulis
// @Accept json
// @Produce json
// @Param id path string true "ID Penulis"
// @Param penulis body model.Penulis true "Data Penulis"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/penulis/{id} [put]
func UpdatePenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var p model.Penulis

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing body",
		})
	}

	err := controller.UpdatePenulisByID(c.Context(), id, p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update penulis: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Penulis berhasil diperbarui",
	})
}

// DeletePenulisByID godoc
// @Summary Hapus penulis
// @Description Menghapus data penulis berdasarkan ID (dengan token Bearer)
// @Tags Penulis
// @Accept json
// @Produce json
// @Param id path string true "ID Penulis"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/penulis/{id} [delete]
func DeletePenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.DeletePenulisByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus penulis: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Penulis dengan ID %s berhasil dihapus", id),
	})
}
