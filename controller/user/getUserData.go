package user

import (
	db "oloapi/database"
	"oloapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&user); res.RowsAffected <= 0 {
		return ctx.JSON(ustatus{StatusCode: errUserNotFound})
	}

	db.DB.Model(&user).Omit("Follows.*").Association("Followers").Count()

	return ctx.JSON(fiber.Map{
		"uuid":            user.UUID,
		"email":           user.Email,
		"name":            user.Name,
		"profile_picture": user.ProfileImage,
		"is_official":     user.IsOfficial,
	})

}
