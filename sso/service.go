package sso

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rohanshrestha09/patra-go/configs"
	"github.com/rohanshrestha09/patra-go/enums"
	"github.com/rohanshrestha09/patra-go/exception"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/modules/user"
	"github.com/rohanshrestha09/patra-go/utils"
)

type SSOService struct {
	userService user.UserService
}

func (s SSOService) HandleLogin(userDetails models.SSOUser) (string, *exception.HttpException) {
	var user models.User

	user.Email = userDetails.Email

	recordExists, err := s.userService.Store.RecordExists(models.User{Email: user.Email})

	if err != nil {
		return "", exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	if !recordExists {
		user.Name = userDetails.Name

		user.Image = userDetails.Image

		user.Password = userDetails.ID

		if err := s.userService.CreateUser(user); err != nil {
			return "", err
		}
	}

	authToken, err := utils.SignJwt(user.Email)

	if err != nil {
		return "", exception.ThrowHttpException(500, err.Error())
	}

	return authToken, nil

}

func getSSOUserInfo[T models.FacebookUser | models.GoogleUser](token, url string) (T, error) {
	var user T

	request, _ := http.NewRequest("GET", url+token, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return user, err
	}

	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&user)

	defer response.Body.Close()

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s SSOService) GetUserDetails(OAuth2Config *configs.OAuth2Config, accessToken string) (models.SSOUser, error) {

	userDetails, err := (func() (models.SSOUser, error) {

		switch OAuth2Config.Provider {
		case enums.Facebook:
			userDetails, err := getSSOUserInfo[models.FacebookUser](accessToken, OAuth2Config.TokenURI)
			return models.SSOUser{
				ID:    userDetails.ID,
				Name:  userDetails.Name,
				Email: userDetails.Email,
				Image: userDetails.Picture.Data.Url,
			}, err

		case enums.Google:
			userDetails, err := getSSOUserInfo[models.GoogleUser](accessToken, OAuth2Config.TokenURI)
			return models.SSOUser{
				ID:    userDetails.ID,
				Name:  userDetails.Name,
				Email: userDetails.Email,
				Image: userDetails.Picture,
			}, err

		default:
			return models.SSOUser{}, errors.New("invalid provider")
		}

	})()

	return userDetails, err
}
