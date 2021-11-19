package main

import (
	"log"

	"oloapi/api/env"
	"oloapi/api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup environmet
	apienv := env.Api{}
	if err := apienv.Setup(); err != nil {
		log.Fatal("Env couldnt load\n", err)
	}

	// CreateServer creates a new Fiber Instance
	app := fiber.New()

	router.SetupRoutes(app, &apienv)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not found"
	})

	log.Fatal(app.Listen(":3001"))
}
