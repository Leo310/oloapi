package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// User represents a User schema
type User struct {
	Base
	Email        string  `json:"email" gorm:"unique; type:varchar; not null"`
	Password     string  `json:"-" gorm:"type:varchar; not null"`
	Name         string  `json:"name" gorm:"type:varchar; not null"`
	ProfileImage string  `json:"profile_image" gorm:"type:varchar; not null"`
	IsVerified   bool    `json:"-" gorm:"default:false; not null"`
	IsOfficial   bool    `json:"is_official" gorm:"default:false; not null"`
	Rating       float32 `json:"rating" gorm:"default: 0.0; type:decimal(1,1);"`
	Follows      []*User `gorm:"many2many:user_relation;foreignKey:UUID;joinForeignKey:follower;References:UUID;joinReferences:following"`
	Followers    []*User `gorm:"many2many:user_relation;foreignKey:UUID;joinForeignKey:following;References:UUID;joinReferences:follower"`
}

// UserErrors represents the error format for user routes
type UserErrors struct {
	Err      bool   `json:"error"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}

func (user *User) AfterCreate(tx *gorm.DB) error {
	tx.Model(user).Update("profile_image", "https://avatars.dicebear.com/api/micah/"+user.Email+".svg")
	return nil
}
