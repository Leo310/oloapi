package models

import (
	"github.com/google/uuid"
)

// User represents a User schema
type User struct {
	Base
	Email        string     `json:"email" gorm:"unique; type:varchar; not null"`
	Password     string     `json:"password" gorm:"type:varchar; not null"`
	Name         string     `json:"name" gorm:"type:varchar; not null"`
	ProfileImage string     `json:"profile_image" gorm:"type:varchar;"`
	IsVerified   bool       `json:"-" gorm:"default:false; not null"`
	IsOfficial   bool       `json:"is_official" gorm:"default:false; not null"`
	Rating       float32    `json:"rating" gorm:"default: 0.0; type:decimal(1,1);"`
	Locations    []Location `gorm:"constraint:OnDelete:CASCADE;"`
}

type Location struct {
	Base
	UserUUID uuid.UUID
	Osm_id   int64  `json:"osm_id"`
	Osm_type string `json:"osm_type"`
}
