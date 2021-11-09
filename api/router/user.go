package router

import (
	"oloapi/api/controller/user"
)

// setupUserRoutes func sets up all the user routes
func setupUserRoutes() {
	userRouter := apirouter.Group("/user")
	// PRIVATE handles all the private user routes that requires authentication
	userPrivRouter := privrouter.Group("/user")

	userRouter.Post("/register", user.RegisterUser)          // Sign Up a user
	userRouter.Post("/login", user.LoginUser)                // Sign In a user
	userRouter.Get("/get-access-token", user.GetAccessToken) // returns a new access_token
	userRouter.Get("/:uuid", user.GetUserData)
	userRouter.Get("/", user.GetUsersData)

	userPrivRouter.Put("/", user.UpdateUser)
	userPrivRouter.Delete("/", user.DeleteUser)
	userPrivRouter.Get("/", user.GetProfileData)
}
