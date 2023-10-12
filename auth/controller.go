package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/enums"
	"github.com/rohanshrestha09/patra-go/interfaces"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/utils"
)

type AuthController struct {
	authService AuthService
}

// Regsiter godoc
//
//	@Summary	Create an account
//	@Tags		Auth
//	@Accept		mpfd
//	@Produce	json
//	@Param		name			formData	string	true	"Name"
//	@Param		email			formData	string	true	"Email"
//	@Param		password		formData	string	true	"Password"			minlength(8)
//	@Param		confirmPassword	formData	string	true	"Confirm Password"	minlength(8)
//	@Param		image			formData	file	false	"File to upload"
//	@Success	201				{object}	dto.ResponseReturn
//	@Router		/auth/register [post]
func (a AuthController) Register(POST interfaces.POST) {

	POST("/auth/register", func(ctx *gin.Context) {

		var registerDto RegisterDto

		if err := ctx.ShouldBindWith(&registerDto, binding.FormMultipart); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		if err := validator.New().Struct(registerDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if file, err := ctx.FormFile("image"); err == nil {
			registerDto.Image, registerDto.ImageName, err = utils.UploadFile(file, enums.USER_DIR, enums.IMAGE)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}

		data, httpError := a.authService.Register(registerDto)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusCreated, dto.GetResponse(map[string]string{"token": data}, "Register Successful"))

	})

}

// Login godoc
//
//	@Summary	Login User
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		body	body		LoginDto	true	"Request Body"
//	@Success	201		{object}	dto.ResponseReturn
//	@Router		/auth/login [post]
func (a AuthController) Login(POST interfaces.POST) {

	POST("/auth/login", func(ctx *gin.Context) {

		var loginDto LoginDto

		if err := ctx.BindJSON(&loginDto); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		data, httpError := a.authService.Login(loginDto)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusCreated, dto.GetResponse(map[string]string{"token": data}, "Login Successful"))

	})

}

// Get Auth Profile godoc
//
//	@Summary	Get profile
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.GetResponseReturn[models.User]
//	@Router		/auth/profile [get]
//	@Security	Bearer
func (a AuthController) GetProfile(GET interfaces.GET) {

	GET("/auth/profile", func(ctx *gin.Context) {

		authUser := ctx.MustGet("authUser").(models.User)

		ctx.JSON(http.StatusOK, dto.GetResponse(authUser, "Profile Fetched"))

	})

}

// Update Auth Profile godoc
//
//	@Summary	Update profile
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	201	{object}	dto.ResponseReturn
//	@Router		/auth/profile [put]
//	@Security	Bearer
func (a AuthController) UpdateProfile(PUT interfaces.PUT) {

	PUT("/auth/profile", func(ctx *gin.Context) {

		authUser := ctx.MustGet("authUser").(models.User)

		var profileUpdateDto ProfileUpdateDTO

		if err := ctx.ShouldBindWith(&profileUpdateDto, binding.FormMultipart); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		if file, err := ctx.FormFile("image"); err == nil {
			profileUpdateDto.ImageUrl, profileUpdateDto.ImageName, err = utils.UploadFile(file, enums.USER_DIR, enums.IMAGE)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}

			if err := utils.DeleteFile(string(enums.USER_DIR) + authUser.ImageName); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}

		httpError := a.authService.UpdateProfile(authUser, profileUpdateDto)

		if httpError != nil {
			ctx.JSON(httpError.Status, dto.Response(httpError.Message))
			return
		}

		ctx.JSON(http.StatusCreated, dto.Response("Profile Updated"))

	})

}
