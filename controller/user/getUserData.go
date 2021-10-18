package user

import (
	"log"
	db "oloapi/database"
	"oloapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(ctx *fiber.Ctx) error {
	// limit := ctx.Query("limit", "10") //default return 10 users

	var users []models.User
	db.DB.Find(&users)
	log.Print(users)

	return ctx.JSON(fiber.Map{
		"hel": "hello",
	})

}
