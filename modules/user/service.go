package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rohanshrestha09/patra-go/dto"
	"github.com/rohanshrestha09/patra-go/enums"
	"github.com/rohanshrestha09/patra-go/exception"
	"github.com/rohanshrestha09/patra-go/models"
	"github.com/rohanshrestha09/patra-go/service"
)

type UserService struct {
	Store service.Store[models.User]
}

func (u UserService) GetUser(args dto.GetArgs[models.User]) (models.User, *exception.HttpException) {
	user, err := u.Store.Get(args)

	if err != nil {
		return user, exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return user, nil
}

func (u UserService) GetUsers(query dto.Query, args dto.GetAllArgs[models.User]) ([]models.User, int64, *exception.HttpException) {
	users, count, err := u.Store.GetAll(query, args)

	if err != nil {
		return users, count, exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return users, count, nil
}

func (u UserService) CreateUser(user models.User) *exception.HttpException {
	err := u.Store.Create(user)

	if err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u UserService) FollowUser(authUser models.User, userId uuid.UUID) *exception.HttpException {
	if err := u.Store.DB.Model(&authUser).Association(enums.FOLLOWING.Value()).Append(&models.User{ID: userId}); err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u UserService) UnfollowUser(authUser models.User, userId uuid.UUID) *exception.HttpException {
	if err := u.Store.DB.Model(&authUser).Association(enums.FOLLOWING.Value()).Delete(&models.User{ID: userId}); err != nil {
		return exception.ThrowHttpException(http.StatusInternalServerError, err.Error())
	}

	return nil
}
