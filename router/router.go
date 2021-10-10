package router

import (
	"github.com/gofiber/fiber/v2"
)

var apirouter fiber.Router

// request go through auth middleware
var privrouter fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	apirouter = app.Group("/api")
	privrouter = apirouter.Group("/privat")

	setupUserRoutes()
}
