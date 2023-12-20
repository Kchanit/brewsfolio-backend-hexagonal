package main

import (
	"log"

	"github.com/Kchanit/brewsfolio-backend/database"
	"github.com/Kchanit/brewsfolio-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:3000",
			AllowCredentials: true,
		}))

	if database.IsUserEmpty() {
		database.SeedUserData()
		database.SeedBeerData()
	}

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":8080")
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
