package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ajm896/adlibai/graph/model"
)

// JournalEntry represents a user journal entry.
type JournalEntry struct {
	gorm.Model
	UserID      uint   `gorm:"not null"` // foreign key to User
	User        User   `gorm:"foreignKey:UserID"`
	Content     string `gorm:"not null"`
	AiGenerated bool   `gorm:"not null"`
}

func (j *JournalEntry) ToQLJournalEntry() *model.JournalEntry {
	return &model.JournalEntry{
		ID:          fmt.Sprint(j.ID),
		Content:     j.Content,
		AiGenerated: j.AiGenerated,
	}
}

func FromQLJournalEntry(journalEntry *model.JournalEntry) *JournalEntry {
	return &JournalEntry{
		//UserID:      journalEntry.UserID,
		Content:     journalEntry.Content,
		AiGenerated: journalEntry.AiGenerated,
	}
}

func (j *JournalEntry) New(content string, aiGenerated bool, user *User) *JournalEntry {
	return &JournalEntry{
		User:        *user,
		Content:     content,
		AiGenerated: aiGenerated,
	}
}
