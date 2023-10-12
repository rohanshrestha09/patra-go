package models

type Chat struct {
	Model
	Users    []*User    `json:"users" gorm:"many2many:user_chats"`
	Messages []*Message `json:"messages"`
}
