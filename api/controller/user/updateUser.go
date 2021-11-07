package user

import (
	"log"
	db "oloapi/api/database"
	"oloapi/api/models"
	"oloapi/api/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// validate if the email, username and password are in correct format
func validateUpdate(user *models.User) ustatus {
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

func UpdateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	var err error
	if user.UUID, err = uuid.Parse(ctx.Locals("uuid").(string)); err != nil {
		// internal server error
		ctx.Status(500)
		return ctx.JSON(ustatus{StatusCode: errServerInternal})
	}

	// first save everything already know about user in user
	db.DB.First(&user)
	notOverridenUUID := user.UUID
	// then overwrite everything
	if err = ctx.BodyParser(user); err != nil {
		return ctx.JSON(ustatus{StatusCode: errReviewInput})
	}
	// user could override a uuid and change the user data off another user
	user.UUID = notOverridenUUID
	// Improvement: only validate changed values
	if status := validateUpdate(user); status.StatusCode != 0 {
		log.Println(status)
		return ctx.JSON(status)
	}
	//user.Locations = make([]models.Location, 0)
	// and update overwritten user
	if dbtx := db.DB.Save(&user); dbtx.Error != nil {
		return ctx.JSON(ustatus{StatusCode: errEmailAlreadyRegistered})
	}
	return ctx.JSON(ustatus{StatusCode: noErr})
}
