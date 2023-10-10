package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/service"
	"gorm.io/gorm"
)

type AuthModule struct {
	Router       *gin.RouterGroup
	Database     *gorm.DB
	UseAuthGuard func() gin.HandlerFunc
}

func (a *AuthModule) Init() {

	userStore := service.Store[models.User]{DB: a.Database}

	userService := user.UserService{Store: userStore}

	authService := AuthService{userService: userService}

	authController := AuthController{authService: authService}

	authController.Register(a.Router.POST)

	authController.Login(a.Router.POST)

	withAuth := a.Router.Use(a.UseAuthGuard())

	authController.GetProfile(withAuth.GET)

	authController.UpdateProfile(withAuth.PUT)
}
