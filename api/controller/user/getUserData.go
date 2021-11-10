package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid") //default return 10 users
	if !validUuid(uuid) {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errReviewInput})
	}

	user := models.User{Locations: []models.Location{}}
	if err := db.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errUserNotFound})
	}
	// TODO better solution with association
	db.DB.Where("user_uuid = ?", uuid).Find(&user.Locations)

	return ctx.JSON(user)

}
