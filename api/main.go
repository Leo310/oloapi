package main

import (
	"log"

	"oloapi/api/database"
	"oloapi/api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// CreateServer creates a new Fiber Instance
func createServer() *fiber.App {
	app := fiber.New()

	return app
}

// Test pipeline report
func main() {
	// TODO why working in olo image? shouldnt because executing oloapi in home directory instead of directory with .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file \n", err)
	}
	// Connect to Postgres
	database.ConnectToDB()
	app := createServer()

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not found"
	})

	log.Fatal(app.Listen(":3001"))
}
