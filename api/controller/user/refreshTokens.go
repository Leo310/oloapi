package user

import (
	"github.com/gofiber/fiber/v2"
)

// RefreshTokens renews expired accesstoken when a valid refresh token is sent
func (userenv *Userenv) RefreshTokens(ctx *fiber.Ctx) error {
	// setting up the authorization cookies
	uuid := ctx.Locals("uuid")
	accessToken, refreshToken := userenv.generateTokens(uuid.(string))
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token_type": "Bearer", "access_token": accessToken, "refresh_token": refreshToken})
}
