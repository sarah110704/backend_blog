package handler

import (
	"Backend/controller"
	"Backend/model"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser godoc
// @Summary Registrasi User Baru
// @Description Menambahkan user baru ke sistem
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.User true "Data User Baru"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/register [post]
func RegisterUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	if err := controller.RegisterUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal registrasi user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User berhasil didaftarkan",
		"data":    user,
	})
}

// LoginUser godoc
// @Summary Login User
// @Description Melakukan login dan menghasilkan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body model.User true "Email dan Password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/login [post]
func LoginUser(c *fiber.Ctx) error {
	var input model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	token, err := controller.LoginUser(c.Context(), input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Email atau password salah",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login berhasil",
		"token":   token,
	})
}
