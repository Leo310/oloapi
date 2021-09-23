package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User represents a User schema
type User struct {
	Base
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// UserErrors represents the error format for user routes
type UserErrors struct {
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}
