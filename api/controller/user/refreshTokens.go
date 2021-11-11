package user

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("PRIV_KEY"))

func RefreshTokens(ctx *fiber.Ctx) error {
	// setting up the authorization cookies
	uuid := ctx.Locals("uuid")
	accessToken, refreshToken := generateTokens(uuid.(string))
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token_type": "Bearer", "access_token": accessToken, "refresh_token": refreshToken})
}
