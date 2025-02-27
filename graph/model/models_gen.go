// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"gorm.io/gorm"
)

type JournalEntry struct {
	gorm.Model
	ID          string `json:"id"`
	User        *User  `json:"user"`
	Content     string `json:"content"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	AiGenerated bool   `json:"aiGenerated"`
}

type Mutation struct {
}

type NotificationPreferences struct {
	gorm.Model
	EmailNotifications bool `json:"emailNotifications"`
	PushNotifications  bool `json:"pushNotifications"`
}

type Query struct {
}

type Streak struct {
	gorm.Model
	User        *User  `json:"user"`
	Count       int32  `json:"count"`
	LastUpdated string `json:"lastUpdated"`
}

type User struct {
	gorm.Model
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserSettings struct {
	gorm.Model
	AiMode                  AIMode                   `json:"aiMode"`
	NotificationPreferences *NotificationPreferences `json:"notificationPreferences"`
}

type AIMode string

const (
	AIModeMimic AIMode = "MIMIC"
	AIModeSilly AIMode = "SILLY"
)

var AllAIMode = []AIMode{
	AIModeMimic,
	AIModeSilly,
}

func (e AIMode) IsValid() bool {
	switch e {
	case AIModeMimic, AIModeSilly:
		return true
	}
	return false
}

func (e AIMode) String() string {
	return string(e)
}

func (e *AIMode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AIMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AIMode", str)
	}
	return nil
}

func (e AIMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
