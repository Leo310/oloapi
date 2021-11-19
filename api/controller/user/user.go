package user

import (
	"time"

	"gorm.io/gorm"
)

type Userenv struct {
	DB                     *gorm.DB
	JwtKey                 []byte
	AccessTokenExpiryTime  time.Duration
	RefreshTokenExpiryTime time.Duration
}
