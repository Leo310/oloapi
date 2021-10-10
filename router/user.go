package router

import (
	"oloapi/user"
	"oloapi/util"
)

// setupUserRoutes func sets up all the user routes
func setupUserRoutes() {
	userRouter := apirouter.Group("/user")
	// PRIVATE handles all the private user routes that requires authentication
	userPrivRouter := privrouter.Group("/user")
	userPrivRouter.Use(util.SecureAuth()) // middleware to secure all routes for this group

	userRouter.Post("/signup", user.CreateUser)              // Sign Up a user
	userRouter.Post("/signin", user.LoginUser)               // Sign In a user
	userRouter.Get("/get-access-token", user.GetAccessToken) // returns a new access_token

	userPrivRouter.Get("/:id?", user.GetUserData)
	userPrivRouter.Get("/follow/:id", user.FollowUser)
}
