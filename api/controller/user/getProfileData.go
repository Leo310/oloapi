package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
)

func GetProfileData(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid")

	user := new(models.User)
	db.DB.Where("uuid = ?", uuid).First(&user)
	// TODO better solution with association
	db.DB.Where("user_uuid = ?", uuid).Find(&user.Locations)

	return ctx.JSON(user)

}
