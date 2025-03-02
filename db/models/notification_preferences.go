package models

import (
	"gorm.io/gorm"
)

// NotificationPreferences holds settings for notifications.
type NotificationPreferences struct {
	gorm.Model
	EmailNotifications bool `gorm:"not null"`
	PushNotifications  bool `gorm:"not null"`
}
