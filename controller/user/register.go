package user

import (
	"log"
	"math/rand"
	db "oloapi/database"
	"oloapi/models"
	"oloapi/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// validate if the email, username and password are in correct format
func validateRegister(user *models.User) ustatus {
	var status ustatus
	if !util.ValidEmail(user.Email) {
		status = ustatus{StatusCode: errEmailInvalid}
	} else if !util.ValidName(user.Name) {
		status = ustatus{StatusCode: errNameInvalid}
	} else if !util.ValidPassword(user.Password) {
		status = ustatus{StatusCode: errPasswordInvalid}
	} else if count := db.DB.Where(&models.User{Email: user.Email}).First(new(models.User)).RowsAffected; count > 0 {
		status = ustatus{StatusCode: errEmailAlreadyRegistered}
	} else {
		status = ustatus{StatusCode: noErr}
	}
	return status
}

// RegisterUser route registers a User into the database
func RegisterUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.JSON(ustatus{StatusCode: errReviewInput})
	}

	if status := validateRegister(user); status.StatusCode != 0 {
		log.Print(status)
		return ctx.JSON(status)
	}

	// TODO Hashing the password with a random salt
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(12-bcrypt.DefaultCost)+bcrypt.DefaultCost,
	)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	// u.ProfileImage = "https://avatars.dicebear.com/api/micah/" + u.Email + ".svg"

	if err := db.DB.Create(&user).Error; err != nil {
		return ctx.JSON(ustatus{StatusCode: errSomeError})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(user.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	ctx.Cookie(accessCookie)
	ctx.Cookie(refreshCookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
