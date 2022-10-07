package repository

import (
	"fiber-repo-pattern/app/model"

	"gorm.io/gorm"
)

type userConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		db: db,
	}
}

func (repo *userConnection) FindAll() []model.User {
	var users []model.User
	repo.db.Find(&users)
	return users
}

func (repo *userConnection) FindById(UserId int) model.User {
	var user model.User
	repo.db.Find(&user, UserId)
	return user
}

func (repo *userConnection) Create(user model.User) model.User {
	repo.db.Create(&user)
	return user
}

func (repo *userConnection) Update(user model.User) model.User {
	repo.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
	return user
}

func (repo *userConnection) Delete(user model.User) error {
	repo.db.Delete(&user)
	return nil
}
