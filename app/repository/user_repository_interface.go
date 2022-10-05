package repository

import "fiber-repo-pattern/app/model"

type UserRepository interface {
	FindAll() []model.User
	FindById(UserId int) model.User
	Create(user model.User) model.User
	Update(user model.User) model.User
	Delete(user model.User) error
}
