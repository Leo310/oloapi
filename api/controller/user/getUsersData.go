package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsersData(ctx *fiber.Ctx) error {
	limitString := ctx.Query("limit", "10") //default return 10 users
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(ustatus{StatusCode: errReviewInput})
	}

	var users []models.User
	db.DB.Limit(limit).Find(&users)

	for index, user := range users {
		// user.Locations = []models.Location{}
		// TODO better solution with association
		db.DB.Where("user_uuid = ?", user.UUID).Find(&user.Locations)
		// because user is only value not reference
		users[index] = user
	}

	return ctx.JSON(users)

}
