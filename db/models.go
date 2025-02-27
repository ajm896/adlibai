package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

// JournalEntry represents a user journal entry.
type JournalEntry struct {
	gorm.Model
	UserID      uint   `gorm:"not null"` // foreign key to User
	User        User   `gorm:"foreignKey:UserID"`
	Content     string `gorm:"not null"`
	AiGenerated bool   `gorm:"not null"`
}

// NotificationPreferences holds settings for notifications.
type NotificationPreferences struct {
	gorm.Model
	EmailNotifications bool `gorm:"not null"`
	PushNotifications  bool `gorm:"not null"`
	// If linked to user settings, you can add a foreign key:
	// UserSettingsID uint
}

// Streak tracks a user's ongoing activity.
type Streak struct {
	gorm.Model
	UserID      uint      `gorm:"not null"` // foreign key to User
	User        User      `gorm:"foreignKey:UserID"`
	Count       int32     `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`
}

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
