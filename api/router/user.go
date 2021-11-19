package router

import (
	"oloapi/api/env"
)

// setupUserRoutes func sets up all the user routes
func setupUserRoutes(apienv *env.API) {
	userRouter := apirouter.Group("/user")
	// PRIVATE handles all the private user routes that requires authentication
	userPrivRouter := privrouter.Group("/user")

	userRouter.Post("/register", apienv.User.RegisterUser) // Sign Up a user
	userRouter.Post("/login", apienv.User.LoginUser)       // Sign In a user
	userRouter.Get("/:uuid", apienv.User.GetUserData)
	userRouter.Get("/", apienv.User.GetUsersData)

	userPrivRouter.Put("/", apienv.User.UpdateUser)
	userPrivRouter.Delete("/", apienv.User.DeleteUser)
	userPrivRouter.Get("/", apienv.User.GetProfileData)
	userPrivRouter.Get("/refreshTokens", apienv.User.RefreshTokens) // returns a new access_token
}
