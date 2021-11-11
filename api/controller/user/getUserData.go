package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type apiUser struct {
	UUID         uuid.UUID `json:"-" gorm:"primaryKey;autoIncrement:false;unique"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	IsOfficial   bool      `json:"is_official"`
	Rating       float32   `json:"rating"`
}

func GetUserData(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid") //default return 10 users
	if !validUuid(uuid) {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errReviewInput})
	}

	user := new(apiUser)
	if err := db.DB.Model(&models.User{}).First(&user, "uuid = ?", uuid).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errUserNotFound})
	}

	return ctx.JSON(user)

}
