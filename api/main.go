package main

import (
	"log"

	"oloapi/api/database"
	"oloapi/api/router"

	"github.com/gofiber/fiber/v2"
)

// CreateServer creates a new Fiber Instance

func createServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	// Connect to Postgres
	database.ConnectToDB()
	app := createServer()

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not found"
	})

	log.Fatal(app.Listen(":3000"))
}
