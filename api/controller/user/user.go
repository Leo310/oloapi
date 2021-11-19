package user

import (
	"time"

	"gorm.io/gorm"
)

// Userenv describes whole user environment
type Userenv struct {
	DB                     *gorm.DB
	JwtKey                 []byte
	AccessTokenExpiryTime  time.Duration
	RefreshTokenExpiryTime time.Duration
}
