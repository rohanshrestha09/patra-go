package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/service"
	"gorm.io/gorm"
)

type UserModule struct {
	Router       *gin.RouterGroup
	Database     *gorm.DB
	UseAuthGuard func() gin.HandlerFunc
}

func (u *UserModule) Init() {

	userStore := service.Store[models.User]{DB: u.Database}

	userService := UserService{Store: userStore}

	userController := UserController{userService: userService}

	userController.GetUser(u.Router.GET)

	withAuth := u.Router.Use(u.UseAuthGuard())

	userController.FollowUser(withAuth.POST)

	userController.UnfollowUser(withAuth.DELETE)

}
