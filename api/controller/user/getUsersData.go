package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type apiUsers struct {
	UUID         uuid.UUID `json:"uuid" gorm:"primaryKey;autoIncrement:false;unique"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	Rating       float32   `json:"rating"`
}

func GetUsersData(ctx *fiber.Ctx) error {
	limitString := ctx.Query("limit", "10") //default return 10 users
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(uerror{ErrorCode: errReviewInput})
	}

	var users []apiUsers
	db.DB.Model(&models.User{}).Limit(limit).Find(&users)

	//dont return locations
	// for index, user := range users {
	// 	// user.Locations = []models.Location{}
	// 	// TODO better solution with association
	// 	db.DB.Where("user_uuid = ?", user.UUID).Find(&user.Locations)
	// 	// because user is only value not reference
	// 	users[index] = user
	// }

	return ctx.JSON(users)

}
