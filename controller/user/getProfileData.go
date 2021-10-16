package user

import (
	db "oloapi/database"
	"oloapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetProfileData(ctx *fiber.Ctx) error {
	id := ctx.Locals("uuid")

	user := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&user); res.RowsAffected <= 0 {
		return ctx.JSON(ustatus{StatusCode: errUserNotFound})
	}

	// TODO return Adress
	return ctx.JSON(fiber.Map{
		"uuid":            user.UUID,
		"email":           user.Email,
		"name":            user.Name,
		"profile_picture": user.ProfileImage,
		"is_official":     user.IsOfficial,
	})

}
