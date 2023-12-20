package utility

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func UserIDFromToken(c *fiber.Ctx) (uint, error) {
	l := c.Locals("user")
	fmt.Println(l)
	if l == nil {
		return 0, fmt.Errorf("User not found in token")
	}

	claims, ok := l.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid token claims type")
	}

	idFloat64, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("Invalid or missing 'id' claim in token")
	}

	id := uint(idFloat64)
	return id, nil
}
