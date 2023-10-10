package user

type CreateUserDto struct {
	Name            string `form:"name" binding:"required"`
	Email           string `form:"email" binding:"required,email" validate:"email"`
	Password        string `form:"password" binding:"required" validate:"gte=8"`
	ConfirmPassword string `form:"confirmPassword" binding:"required" validate:"eqfield=Password"`
	Image           string `json:"image"`
	ImageName       string `json:"imageName"`
}
