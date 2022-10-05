package service

import (
	"fiber-repo-pattern/app/model"
	"fiber-repo-pattern/app/request"
)

type UserService interface {
	FindAll() []model.User
	FindById(UserId int) model.User
	Create(dataUser request.UserCreateValidation) model.User
	Update(dataUser request.UserUpdateValidation) model.User
	Delete(dataUser model.User) error
}
