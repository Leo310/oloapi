package user

import (
	"log"
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid") //default return 10 users
	if !validUuid(uuid) {
		log.Println(errReviewInput)
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(ustatus{StatusCode: errReviewInput})
	}

	user := models.User{Locations: []models.Location{}}
	if err := db.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		log.Println(errUserNotFound)
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(ustatus{StatusCode: errUserNotFound})
	}
	// TODO better solution with association
	db.DB.Where("user_uuid = ?", uuid).Find(&user.Locations)

	return ctx.JSON(user)

}
