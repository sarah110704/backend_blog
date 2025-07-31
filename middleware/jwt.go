package middleware

import (
	"Backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
<<<<<<< HEAD
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
=======
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("RAHASIA_LO_JANGAN_SEBARIN"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
>>>>>>> 36605f5109d3743dcee478d3c815d3b15f6f91d5
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Check if it's a Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Bearer token required",
			})
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is empty",
			})
		}

		// Validate token using our utils function
		claims, err := utils.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
				"debug": err.Error(),
			})
		}

		// Store user info in context
		c.Locals("user_claims", claims)

		return c.Next()
	}
}
