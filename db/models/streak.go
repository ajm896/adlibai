package models

import (
	"time"

	"gorm.io/gorm"
)

// Streak tracks a user's ongoing activity.
type Streak struct {
	gorm.Model
	UserID      uint      `gorm:"not null"` // foreign key to User
	User        User      `gorm:"foreignKey:UserID"`
	Count       int32     `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`
}
