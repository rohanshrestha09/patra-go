package user

import (
	"net/http"

	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/exception"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/service"
)

type UserService struct {
	Store service.Store[models.User]
}

func (u UserService) GetUser(email string) (models.User, *exception.HttpException) {

	user, err := u.Store.Get(dto.GetArgs[models.User]{
		Filter:  models.User{Email: email},
		Exclude: []string{"Email"},
	})

	if err != nil {
		return models.User{}, exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return user, nil

}

func (u UserService) CreateUser(user models.User) *exception.HttpException {
	err := u.Store.Create(user)

	if err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u UserService) FollowUser(authUser models.User, userId string) *exception.HttpException {
	err := u.Store.DB.Create(&models.UserFollows{
		FollowedByID: authUser.ID,
		FollowingID:  userId,
	}).Error

	if err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u UserService) UnfollowUser(authUser models.User, userId string) *exception.HttpException {
	err := u.Store.DB.Unscoped().Delete(&models.UserFollows{
		FollowedByID: authUser.ID,
		FollowingID:  userId,
	}).Error

	if err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}
