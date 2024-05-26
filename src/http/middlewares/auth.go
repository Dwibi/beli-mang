package middlewares

import (
	// "github.com/gofiber/fiber"

	"os"
	"strings"

	"github.com/Dwibi/beli-mang/src/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized!")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &helpers.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized!")
	}

	claims, ok := token.Claims.(*helpers.Claims)

	if !ok || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized!")
	}

	c.Locals("userId", claims.UserId)

	return c.Next()
}
