package env

import (
	"errors"
	"oloapi/api/controller/user"
	"oloapi/api/database"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// TODO implement interface methods
type Api struct {
	User interface {
		Authenticator() func(*fiber.Ctx) error
		RegisterUser(*fiber.Ctx) error
		LoginUser(*fiber.Ctx) error
		UpdateUser(*fiber.Ctx) error
		DeleteUser(*fiber.Ctx) error
		GetProfileData(*fiber.Ctx) error
		GetUserData(*fiber.Ctx) error
		GetUsersData(*fiber.Ctx) error
		RefreshTokens(*fiber.Ctx) error
	}
	Item interface {
	}
	Database interface {
		Connect() *gorm.DB
	}
}

func (env *Api) Setup() error {
	// TODO why working in olo image? shouldnt because executing oloapi in home directory instead of directory with .env file
	if err := godotenv.Load(); err != nil {
		return errors.New("error loading env file")
	}
	env.Database = &database.DBenv{
		Ip:       os.Getenv("PSQL_IP"),
		Port:     os.Getenv("PSQL_PORT"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASS"),
		Name:     os.Getenv("PSQL_DBNAME"),
	}
	env.User = &user.Userenv{
		// Connect to Postgres
		DB:                     env.Database.Connect(),
		JwtKey:                 []byte(os.Getenv("PRIV_KEY")),
		AccessTokenExpiryTime:  15 * time.Minute,
		RefreshTokenExpiryTime: 30 * 24 * time.Hour,
	}
	return nil
}
