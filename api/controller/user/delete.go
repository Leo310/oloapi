package user

import (
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser currently loged in user
func (userenv *Userenv) DeleteUser(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid")
	userenv.DB.Where("uuid = ?", uuid).Delete(&models.User{})
	return ctx.Status(fiber.StatusOK).Send(nil)
}
