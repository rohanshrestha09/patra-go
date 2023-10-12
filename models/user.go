package models

import (
	"github.com/google/uuid"
	"github.com/rohanshrestha09/patra-go/enums"
)

type User struct {
	Model
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	Email     string         `json:"email,omitempty" gorm:"not null;unique"`
	Name      string         `json:"name" gorm:"not null"`
	Bio       string         `json:"bio" gorm:"type:text"`
	Password  string         `json:"-" gorm:"not null"`
	Image     string         `json:"image"`
	ImageName string         `json:"imageName"`
	Provider  enums.Provider `json:"provider" gorm:"type:provider;default:EMAIL;not null"`
	Following []*User        `json:"following" gorm:"many2many:user_follows"`
	Chats     []*Chat        `json:"chats" gorm:"many2many:user_chats"`
	Messages  []*Message     `json:"messages" gorm:"foreignKey:SenderID"`
}
