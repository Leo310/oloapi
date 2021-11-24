package env

import (
	"errors"
	"oloapi/api/controller/user"
	"oloapi/api/database"
	"os"
	"path"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// API contains whole environment as interfaces
type API struct {
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

// Setup initializes whole enviornment
func (env *API) Setup() error {
	// checks which env file there is in env folder and laods it
	envPath := path.Join(os.Getenv("GOPATH"), "src/oloapi/api/env/.env")
	if err := godotenv.Load(envPath + "local"); err != nil {
		if err = godotenv.Load(envPath + "deploy"); err != nil {
			if err = godotenv.Load(envPath + "stage"); err != nil {
				return errors.New("no env file found (supported: .envlocal, .envstage, .envdeploy)")
			}
		}
	}
	env.Database = &database.DBenv{
		IP:       os.Getenv("PSQL_IP"),
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
