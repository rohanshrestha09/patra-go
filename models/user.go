package models

import "github.com/rohanshrestha09/patra-go/enums"

type User struct {
	Model
	ID        string         `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	Email     string         `json:"email,omitempty" gorm:"not null;unique"`
	Name      string         `json:"name" gorm:"not null"`
	Bio       string         `json:"bio" gorm:"type:text"`
	Password  string         `json:"-" gorm:"not null"`
	Image     string         `json:"image"`
	ImageName string         `json:"imageName"`
	Provider  enums.Provider `json:"provider" gorm:"type:provider;default:EMAIL;not null"`
}

type UserFollows struct {
	Model
	FollowedByID string `json:"followedById" gorm:"type:uuid;not null"`
	FollowedBy   *User  `json:"followedBy"`
	FollowingID  string `json:"followingId" gorm:"type:uuid;not null"`
	Following    *User  `json:"following"`
}
