package user

import (
	"log"
	db "oloapi/database"
	"oloapi/models"
	"oloapi/util"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid") //default return 10 users
	if !util.ValidUuid(uuid) {
		log.Println(errReviewInput)
		ctx.Status(400)
		return ctx.JSON(ustatus{StatusCode: errReviewInput})
	}

	user := new(models.User)
	if err := db.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		log.Println(errUserNotFound)
		ctx.Status(400)
		return ctx.JSON(ustatus{StatusCode: errUserNotFound})
	}

	return ctx.JSON(user)

}
