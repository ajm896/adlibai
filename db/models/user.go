package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ajm896/adlibai/graph/model"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

func (u *User) ToQLUser() *model.User {
	return &model.User{
		ID:       fmt.Sprint(u.ID),
		Username: u.Username,
		Email:    u.Email,
	}
}

func FromQLUser(user *model.User) *User {
	return NewUser(user.Username, user.Email)
}

func NewUser(username, email string) *User {
	return &User{
		Username: username,
		Email:    email,
	}
}
