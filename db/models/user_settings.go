package models

import (
	"gorm.io/gorm"
)

// UserSettings stores the user's configuration including AI mode and notification preferences.
type UserSettings struct {
	gorm.Model
	UserID uint `gorm:"not null;unique"` // one-to-one relation with User
	User   User `gorm:"foreignKey:UserID"`
	// Storing AIMode as string, matching the generated GraphQL enum.
	AiMode string `gorm:"not null"`
	// Association to NotificationPreferences. Uncomment and adjust if needed.
	// NotificationPreferencesID uint
	// NotificationPreferences   NotificationPreferences `gorm:"foreignKey:NotificationPreferencesID"`
}
