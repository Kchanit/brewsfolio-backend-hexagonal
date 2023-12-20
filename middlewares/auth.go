package middlewares

import (
	"fmt"
	"os"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SCERET_KEY"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func AuthMiddleware(c *fiber.Ctx) error {
	fmt.Println("AuthMiddleware")
	tokenString := extractTokenFromHeader(c)

	if tokenString == "" {
		return c.Next()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the token is signed with the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Parse the claims
		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("Token claims:", claims)
		// Store user ID in locals
		c.Locals("user", claims)

		// Return the secret key
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	if !token.Valid {
		fmt.Println("Invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	fmt.Println("AdminMiddleware")
	tokenString := extractTokenFromHeader(c)
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Get the user from the locals
	user := c.Locals("user").(jwt.MapClaims)
	role := user["role"]

	// Check if the user is an admin
	if role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

func extractTokenFromHeader(c *fiber.Ctx) string {
	// Get the Authorization header value
	authHeader := c.Get("Authorization")

	// Check if the header is not present
	if authHeader == "" {
		return ""
	}

	// Split the header value to get the token part
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	// Return the token part
	return parts[1]
}
