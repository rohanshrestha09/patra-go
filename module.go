package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/auth"
	"github.com/rohanshrestha09/patra-go/configs"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/service"
	"github.com/rohanshrestha09/patra-go/sso"
)

func InitModule(router *gin.RouterGroup) {

	database := configs.InitializeDatabase()

	userStore := service.Store[models.User]{DB: database}

	userService := user.UserService{Store: userStore}

	authGuard := auth.AuthGuard{UserService: userService}

	authModule := auth.AuthModule{
		Router:       router,
		Database:     database,
		UseAuthGuard: authGuard.UseAuthGuard,
	}

	authModule.Init()

	ssoModule := sso.SSOModule{
		Router:   router,
		Database: database,
	}

	ssoModule.Init()

	userModule := user.UserModule{
		Router:       router,
		Database:     database,
		UseAuthGuard: authGuard.UseAuthGuard,
	}

	userModule.Init()

}
