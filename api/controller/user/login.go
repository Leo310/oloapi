package user

import (
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser user signin
func (userenv *Userenv) LoginUser(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.JSON(uerror{ErrorCode: errReviewInput})
	}

	// check if a user exists
	user := new(models.User)
	if res := userenv.DB.Where(
		&models.User{Email: input.Email}).First(&user); res.RowsAffected <= 0 {
		return ctx.Status(fiber.StatusTeapot).JSON(uerror{ErrorCode: errCredentialsInvalid})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password+user.Salt)); err != nil {
		return ctx.Status(fiber.StatusTeapot).JSON(uerror{ErrorCode: errCredentialsInvalid})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := userenv.generateTokens(user.UUID.String())
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token_type": "Bearer", "access_token": accessToken, "refresh_token": refreshToken})
}
