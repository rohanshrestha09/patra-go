package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/interfaces"
	"github.com/rohanshrestha09/patra-go/models"
)

type UserController struct {
	userService UserService
}

// Get User godoc
//
//	@Summary	Get a user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		email	path		string	true	"Email"
//	@Success	200		{object}	dto.GetResponseReturn[models.User]
//	@Router		/user/{email} [get]
func (u UserController) GetUser(GET interfaces.GET) {

	GET("/user/:email", func(ctx *gin.Context) {

		data, httpError := u.userService.GetUser(ctx.Param("email"))

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusOK, dto.GetResponse(data, "User Fetched"))

	})

}

func (u UserController) FollowUser(POST interfaces.POST) {

	POST("/user/:id/follow", func(ctx *gin.Context) {

		authUser := ctx.MustGet("auth").(models.User)

		userId := ctx.Param("id")

		httpError := u.userService.FollowUser(authUser, userId)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusOK, dto.Response("Followed"))

	})

}

func (u UserController) UnfollowUser(DELETE interfaces.DELETE) {

	DELETE("/user/:id/follow", func(ctx *gin.Context) {

		authUser := ctx.MustGet("auth").(models.User)

		userId := ctx.Param("id")

		httpError := u.userService.UnfollowUser(authUser, userId)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusOK, dto.Response("Unfollowed"))

	})

}
