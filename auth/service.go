package auth

import (
	"net/http"

	"github.com/rohanshrestha09/patra-go/exception"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService user.UserService
}

func (a AuthService) Register(registerDto RegisterDto) (string, *exception.HttpException) {

	user := models.User{
		Name:      registerDto.Name,
		Email:     registerDto.Email,
		Password:  registerDto.Password,
		Image:     registerDto.Image,
		ImageName: registerDto.ImageName,
	}

	recordExists, err := a.userService.Store.RecordExists(models.User{Email: user.Email})

	if err != nil {
		return "", exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	if recordExists {
		return "", exception.ThrowHttpException(http.StatusForbidden, "User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", exception.ThrowHttpException(http.StatusUnprocessableEntity, "Something went wrong")
	}

	user.Password = string(hash)

	if err := a.userService.CreateUser(user); err != nil {
		return "", err
	}

	authToken, err := utils.SignJwt(user.Email)

	if err != nil {
		return "", exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return authToken, nil

}

func (a AuthService) Login(loginDto LoginDto) (string, *exception.HttpException) {

	user, httpError := a.userService.GetUser(loginDto.Email)

	if httpError != nil {
		return "", httpError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", exception.ThrowHttpException(http.StatusUnauthorized, "Incorrect Password")
	}

	authToken, err := utils.SignJwt(user.Email)

	if err != nil {
		return "", exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return authToken, nil

}

func (a AuthService) UpdateProfile(authUser models.User, profileUpdateDto ProfileUpdateDTO) *exception.HttpException {

	data := models.User{
		Name:      profileUpdateDto.Name,
		Bio:       profileUpdateDto.Bio,
		Image:     profileUpdateDto.ImageUrl,
		ImageName: profileUpdateDto.ImageName,
	}

	err := a.userService.Store.Update(authUser, data)

	if err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil

}
