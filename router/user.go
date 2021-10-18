package router

import (
	"oloapi/controller/user"
	"oloapi/util"
)

// setupUserRoutes func sets up all the user routes
func setupUserRoutes() {
	userRouter := apirouter.Group("/user")
	// PRIVATE handles all the private user routes that requires authentication
	userPrivRouter := privrouter.Group("/user")
	userPrivRouter.Use(util.SecureAuth()) // middleware to secure all routes for this group

	userRouter.Post("/register", user.RegisterUser)          // Sign Up a user
	userRouter.Post("/login", user.LoginUser)                // Sign In a user
	userRouter.Get("/get-access-token", user.GetAccessToken) // returns a new access_token
	userRouter.Get("/:uuid", user.GetUserData)
	userRouter.Get("/", user.GetUsersData)

	userPrivRouter.Delete("/delete", user.DeleteUser)
	userPrivRouter.Get("/", user.GetProfileData)
}
