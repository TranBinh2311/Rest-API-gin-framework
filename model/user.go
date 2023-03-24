package model

import (
	"time"
)

type User struct {
	ID                int    `gorm:"primaryKey;not null"`
	Username          string `gorm:"not null"`
	Password          string `gorm:"not null"`
	FullName          string `gorm:"not null"`
	Email             string `gorm:"unique;not null"`
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}
