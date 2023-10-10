package sso

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/configs"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/interfaces"
)

type SSOController struct {
	ssoService SSOService
}

func (s SSOController) InitFacebookLogin(GET interfaces.GET) {
	GET("/login/facebook", func(ctx *gin.Context) {

		state, err := configs.GetRandomOAuthStateString()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		session := sessions.Default(ctx)

		session.Set("state", state)

		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		url := configs.GetFacebookOAuthConfig().AuthCodeURL(state)

		ctx.Redirect(http.StatusTemporaryRedirect, url)

	})
}

func (s SSOController) HandleFacebookLogin(GET interfaces.GET) {

	GET("/facebook/callback", func(ctx *gin.Context) {

		session := sessions.Default(ctx)

		if ctx.Query("state") != session.Get("state") {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		OAuth2Config := configs.GetFacebookOAuthConfig()

		token, err := OAuth2Config.Exchange(context.TODO(), ctx.Query("code"))

		if err != nil || token == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		userDetails, err := s.ssoService.GetUserDetails(OAuth2Config, token.AccessToken)

		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		authToken, httpError := s.ssoService.HandleLogin(userDetails)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.SetCookie("token", authToken, 30*1400*60, "/", "localhost", false, true)

		ctx.Redirect(http.StatusTemporaryRedirect, "/profile")

	})

}

func (s SSOController) InitGoogleLogin(GET interfaces.GET) {
	GET("/login/google", func(ctx *gin.Context) {

		state, err := configs.GetRandomOAuthStateString()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		session := sessions.Default(ctx)

		session.Set("state", state)

		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		url := configs.GetGoogleOAuthConfig().AuthCodeURL(state)

		ctx.Redirect(http.StatusTemporaryRedirect, url)

	})
}

func (s SSOController) HandleGoogleLogin(GET interfaces.GET) {

	GET("/google/callback", func(ctx *gin.Context) {

		session := sessions.Default(ctx)

		if ctx.Query("state") != session.Get("state") {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		OAuth2Config := configs.GetGoogleOAuthConfig()

		token, err := OAuth2Config.Exchange(context.TODO(), ctx.Query("code"))

		if err != nil || token == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		userDetails, err := s.ssoService.GetUserDetails(OAuth2Config, token.AccessToken)

		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		authToken, httpError := s.ssoService.HandleLogin(userDetails)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.SetCookie("token", authToken, 30*1400*60, "/", "localhost", false, true)

		ctx.Redirect(http.StatusTemporaryRedirect, "/profile")

	})

}
