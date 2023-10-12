package models

import "github.com/google/uuid"

type Message struct {
	Model
	Content    string    `json:"content" gorm:"type:text"`
	Attachment string    `json:"attachment"`
	SenderID   uuid.UUID `json:"senderId" gorm:"type:uuid;not null"`
	Sender     *User     `json:"sender" gorm:"not null"`
	ChatID     uint      `json:"chatId" gorm:"not null"`
	Chat       *Chat     `json:"chat"`
}
