package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
	// and update overwritten user
	if dbtx := db.DB.Save(&user); dbtx.Error != nil {
		return ctx.JSON(ustatus{StatusCode: errEmailAlreadyRegistered})
	}
	return ctx.JSON(ustatus{StatusCode: noErr})
}
