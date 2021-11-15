package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type apiProfile struct {
	UUID         uuid.UUID     `json:"-" gorm:"primaryKey;autoIncrement:false;unique"`
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	ProfileImage string        `json:"profile_image"`
	IsOfficial   bool          `json:"is_official"`
	Rating       float32       `json:"rating"`
	Locations    []apiLocation `json:"locations" gorm:"foreignKey:UserUUID"`
}

type apiLocation struct {
	UserUUID uuid.UUID `json:"-"`
	OsmID    int64     `json:"osm_id"`
	OsmType  string    `json:"osm_type"`
}

// GetProfileData return data of an user who is logged in
func GetProfileData(ctx *fiber.Ctx) error {
	uuid := ctx.Locals("uuid")

	user := new(apiProfile)
	// could be cleaner with preload but doesnt work because of bug https://github.com/go-gorm/gorm/issues/4015
	db.DB.Model(&models.User{}).Where("uuid = ?", uuid).First(user)
	// TODO better solution with association
	db.DB.Model(&models.Location{}).Where("user_uuid = ?", uuid).Find(&user.Locations)

	return ctx.JSON(user)
}
