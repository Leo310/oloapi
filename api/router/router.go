package router

import (
	"oloapi/api/middleware"

	"github.com/gofiber/fiber/v2"
)

var apirouter fiber.Router

// request go through auth middleware
var privrouter fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	apirouter = app.Group("/api")
	privrouter = apirouter.Group("/private")
	privrouter.Use(middleware.SecureAuth()) // middleware to secure all routes for this group

	setupUserRoutes()
}
