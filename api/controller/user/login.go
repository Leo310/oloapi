package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"
	"oloapi/api/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(ustatus{StatusCode: errReviewInput})
	}

	// check if a user exists
	u := new(models.User)
	if res := db.DB.Where(
		&models.User{Email: input.Email}).First(&u); res.RowsAffected <= 0 {
		return c.JSON(ustatus{StatusCode: errCredentialsInvalid})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return c.JSON(ustatus{StatusCode: errCredentialsInvalid})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
