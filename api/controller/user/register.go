package user

import (
	"math/rand"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// validate if the email, username and password are in correct format
func validateRegister(user *models.User) errorCode {
	var error errorCode
	if !validEmail(user.Email) {
		error = errEmailInvalid
	} else if !validName(user.Name) {
		error = errNameInvalid
	} else if !validPassword(user.Password) {
		error = errPasswordInvalid
		//only check first address because client only sends one location on register
	} else if _, err := getValidLookup(user.Locations[0].OsmID, user.Locations[0].OsmType); err != nil {
		error = errLocationNotFound
	} else {
		error = "NO_ERROR"
	}
	return error
}

func generateRandomSalt() string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$%?"
	var salt string
	for i := 0; i < 8; i++ {
		salt += string(characters[rand.Intn(len(characters))])
	}
	return salt
}

// RegisterUser sign up a user
func (userenv *Userenv) RegisterUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errReviewInput})
	}
	if error := validateRegister(user); error != "NO_ERROR" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: error})
	}
	if count := userenv.DB.Where(&models.User{Email: user.Email}).First(new(models.User)).RowsAffected; count > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: errEmailAlreadyRegistered})
	}
	//user.Locations = make([]models.Location, 0)
	//for _, reqLocation := range user.Locations {
	//	if possibleLocations, err := GetValidAddress(reqLocation.Street, reqLocation.Housenumber, reqLocation.City); err == nil {
	//		//only use first returned Address of Geoapi because it has highest possibility to be the right address
	//		location := models.Location{Osm_id: possibleLocations[0].Osm_id, Osm_type: possibleLocations[0].Osm_type}
	//		user.Locations = append(user.Locations, location)
	//	} else {
	//		return ctx.JSON(ustatus{StatusCode: errLocationNotFound})
	//	}
	//}

	user.Salt = generateRandomSalt()
	password := []byte(user.Password + user.Salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(12-bcrypt.DefaultCost)+bcrypt.DefaultCost,
	)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	// u.ProfileImage = "https://avatars.dicebear.com/api/micah/" + u.Email + ".svg"
	user.ProfileImage = "https://upload.wikimedia.org/wikipedia/commons/8/89/Portrait_Placeholder.png"

	if err := userenv.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: errSomeError})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := userenv.generateTokens(user.UUID.String())
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token_type": "Bearer", "access_token": accessToken, "refresh_token": refreshToken})
}
