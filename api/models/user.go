package models

import (
	"github.com/google/uuid"
)

// User represents a User schema
type User struct {
	Base
	Email        string     `json:"email" gorm:"unique; type:varchar; not null"`
	Password     string     `json:"password" gorm:"type:varchar; not null"`
	Salt         string     `json:"salt" gorm:"not null"`
	Name         string     `json:"name" gorm:"type:varchar; not null"`
	ProfileImage string     `json:"profile_image" gorm:"type:varchar;"`
	IsVerified   bool       `json:"-" gorm:"default:false; not null"`
	IsOfficial   bool       `json:"is_official" gorm:"default:false; not null"`
	Rating       float32    `json:"rating" gorm:"default: 0.0; type:decimal(1,1);"`
	Locations    []Location `json:"locations" gorm:"constraint:OnDelete:CASCADE;"`
}

// Location represents a Location schema
type Location struct {
	Base
	UserUUID uuid.UUID
	OsmID    int64  `json:"osm_id"`
	OsmType  string `json:"osm_type"`
}
