package model

import (
	"time"

	"gorm.io/gorm"
)

// User model - `users` table
type User struct {
	UserID    uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName string         `json:"FirstName,omitempty"`
	LastName  string         `json:"LastName,omitempty"`
	IDAuth    uint64         `json:"-"`
}
