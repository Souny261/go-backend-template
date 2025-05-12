package domain

import (
	"time"

	"gorm.io/gorm"
)

// User is the gorm model for the users table
type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(255);uniqueIndex"`
	Username     string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash []byte
	Name         string `gorm:"type:varchar(255);not null"`
	Avatar       string `gorm:"type:varchar(255)"`
	Phone        string `gorm:"type:varchar(20)"`
	Verified     bool   `gorm:"type:tinyint(1);not null;default:0"`
	Status       bool   `gorm:"type:tinyint(1);not null;default:1"`
	LastActive   time.Time
}
