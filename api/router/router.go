package router

import (
	"oloapi/api/env"

	"github.com/gofiber/fiber/v2"
)

var apirouter fiber.Router

// request go through auth middleware
var privrouter fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App, apienv *env.Api) {
	apirouter = app.Group("/api")
	privrouter = apirouter.Group("/private")
	privrouter.Use(apienv.User.Authenticator()) // middleware to secure all routes for this group

	setupUserRoutes(apienv)
}
