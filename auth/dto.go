package auth

import (
	"mime/multipart"
)

type RegisterDto struct {
	Name            string `form:"name" binding:"required"`
	Email           string `form:"email" binding:"required,email" validate:"email"`
	Password        string `form:"password" binding:"required" validate:"gte=8"`
	ConfirmPassword string `form:"confirmPassword" binding:"required" validate:"eqfield=Password"`
	Image           string `json:"image"`
	ImageName       string `json:"imageName"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileUpdateDTO struct {
	Name      string                `form:"name"`
	Bio       string                `form:"bio"`
	Image     *multipart.FileHeader `form:"image"`
	ImageUrl  string
	ImageName string
}
