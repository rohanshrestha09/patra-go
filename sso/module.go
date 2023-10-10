package sso

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/service"
	"gorm.io/gorm"
)

type SSOModule struct {
	Router   *gin.RouterGroup
	Database *gorm.DB
}

func (s *SSOModule) Init() {

	userStore := service.Store[models.User]{DB: s.Database}

	userService := user.UserService{Store: userStore}

	ssoService := SSOService{userService: userService}

	ssoController := SSOController{ssoService: ssoService}

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	s.Router.Use(sessions.Sessions("auth-session", store))

	ssoController.InitFacebookLogin(s.Router.GET)

	ssoController.HandleFacebookLogin(s.Router.GET)

	ssoController.InitGoogleLogin(s.Router.GET)

	ssoController.HandleGoogleLogin(s.Router.GET)

}
