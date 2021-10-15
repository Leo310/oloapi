package user

import (
	db "oloapi/database"
	"oloapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) error {
	id := c.Locals("id")
	if c.Params("id") != "" {
		id = c.Params("id")
	}

	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
		return c.JSON(ustatus{StatusCode: errUserNotFound})
	}

	db.DB.Model(&u).Omit("Follows.*").Association("Followers").Count()

	return c.JSON(fiber.Map{
		"uuid":            u.UUID,
		"email":           u.Email,
		"name":            u.Name,
		"profile_picture": u.ProfileImage,
		"is_official":     u.IsOfficial,
	})

}
