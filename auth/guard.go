package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/utils"
)

type AuthGuard struct {
	UserService user.UserService
}

func (a AuthGuard) UseAuthGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var jwtToken string

		bearerToken := ctx.GetHeader("Authorization")

		if strings.HasPrefix(bearerToken, "Bearer") && len(strings.Split(bearerToken, " ")) == 2 {
			jwtToken = strings.Split(bearerToken, " ")[1]
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response("Unauthorized"))
			return
		}

		claims, token, err := utils.ParseJwt(jwtToken)

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response("Unauthorized"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response(err.Error()))
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response("Unauthorized"))
			return
		}

		args := dto.GetArgs[models.User]{
			Filter: models.User{Email: claims.Email},
		}

		data, httpError := a.UserService.GetUser(args)

		if httpError != nil {
			ctx.AbortWithStatusJSON(httpError.Status, dto.Response(err.Error()))
			return
		}

		ctx.Set("authUser", data)

		ctx.Next()

	}
}
