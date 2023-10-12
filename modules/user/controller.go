package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/enums"
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

		args := dto.GetArgs[models.User]{
			Filter:  models.User{Email: ctx.Param("email")},
			Exclude: []string{"Email"},
		}

		data, httpError := u.userService.GetUser(args)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusOK, dto.GetResponse(data, "User Fetched"))

	})

}

// Get All User godoc
//
//	@Summary	Get all user
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Param		search	query		string	false	"Search"
//	@Success	200		{object}	dto.GetAllResponse[models.User]
//	@Router		/user/ [get]
func (u UserController) GetUsers(GET interfaces.GET) {

	GET("/user", func(ctx *gin.Context) {

		var query dto.Query

		if err := ctx.BindQuery(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, dto.Response(err.Error()))
		}

		args := dto.GetAllArgs[models.User]{
			Search: map[enums.SearchColumn]string{
				enums.NAME_COLUMN:  ctx.Query("search"),
				enums.EMAIL_COLUMN: ctx.Query("search"),
			},
		}

		data, count, httpError := u.userService.GetUsers(query, args)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		pagination := dto.Pagination{
			Page:  query.Page,
			Size:  query.Size,
			Count: count,
		}

		currentPage, totalPage := pagination.GetPages()

		ctx.JSON(http.StatusOK, dto.GetAllResponse[models.User]{
			Message:     "Users fetched",
			Data:        data,
			Count:       count,
			CurrentPage: currentPage,
			TotalPage:   totalPage,
		})

	})

}

// Follow User godoc
//
//	@Summary	Follow User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User Id"
//	@Success	201	{object}	dto.ResponseReturn
//	@Router		/user/{id}/follow [post]
//	@Security	Bearer
func (u UserController) FollowUser(POST interfaces.POST) {

	POST("/user/:id/follow", func(ctx *gin.Context) {

		authUser := ctx.MustGet("authUser").(models.User)

		userId, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.Response(err.Error()))
			return
		}

		httpError := u.userService.FollowUser(authUser, userId)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusCreated, dto.Response("Followed"))

	})

}

// Unfollow User godoc
//
//	@Summary	Unfollow User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User Id"
//	@Success	201	{object}	dto.ResponseReturn
//	@Router		/user/{id}/follow [delete]
//	@Security	Bearer
func (u UserController) UnfollowUser(DELETE interfaces.DELETE) {

	DELETE("/user/:id/follow", func(ctx *gin.Context) {

		authUser := ctx.MustGet("authUser").(models.User)

		userId, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.Response(err.Error()))
			return
		}

		httpError := u.userService.UnfollowUser(authUser, userId)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusCreated, dto.Response("Unfollowed"))

	})

}
