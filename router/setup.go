package router

import (
	"oloapi/util"

	"github.com/gofiber/fiber/v2"
)

// USER handles all the user routes
var USER fiber.Router
var PRIVATE fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// PRIVATE handles all the private user routes that requires authentication
	PRIVATE = api.Group("/private")
	PRIVATE.Use(util.SecureAuth()) // middleware to secure all routes for this group

	USER = api.Group("/user")
	SetupUserRoutes()
}
