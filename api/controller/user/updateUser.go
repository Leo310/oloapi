package user

import (
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// validate if the email, username and password are in correct format
func validateUpdate(user *models.User) errorCode {
	var error errorCode
	if !validEmail(user.Email) {
		error = errEmailInvalid
	} else if !validName(user.Name) {
		error = errNameInvalid
	} else if !validPassword(user.Password) {
		error = errPasswordInvalid
	} else {
		for i := range user.Locations {
			if _, err := getValidLookup(user.Locations[i].OsmID, user.Locations[i].OsmType); err != nil {
				error = errLocationNotFound
				return error
			}
		}
		error = "NO_ERROR"
	}
	return error
}

// UpdateUser updates User data of authorized user
func (userenv *Userenv) UpdateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	var err error
	if user.UUID, err = uuid.Parse(ctx.Locals("uuid").(string)); err != nil {
		// internal server error
		return ctx.Status(fiber.StatusInternalServerError).JSON(uerror{ErrorCode: errServerInternal})
	}

	// first save everything already know about user in user
	userenv.DB.First(&user)
	notOverridenUUID := user.UUID
	// then overwrite everything
	if err = ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errReviewInput})
	}
	// user could override a uuid and change the user data off another user
	user.UUID = notOverridenUUID

	// Improvement: only validate changed values
	if error := validateUpdate(user); error != "NO_ERROR" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: error})
	}
	//check if mail already exists
	if count := userenv.DB.Where(&models.User{Email: user.Email}).First(new(models.User)).RowsAffected; count > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: errEmailAlreadyRegistered})
	}

	//user.Locations = make([]models.Location, 0)
	// and update overwritten user
	if dbtx := userenv.DB.Save(&user); dbtx.Error != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(uerror{ErrorCode: errEmailAlreadyRegistered})
	}
	return ctx.Status(fiber.StatusOK).Send(nil)
}
