package user

import (
	db "oloapi/database"
	"oloapi/models"
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

	return ctx.JSON(users)

}
