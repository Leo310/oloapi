package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid")
	db.DB.Where("uuid = ?", uuid).Delete(&models.User{})
	return ctx.JSON(ustatus{StatusCode: noErr})
}
