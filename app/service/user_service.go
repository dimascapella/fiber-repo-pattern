package service

import (
	"fiber-repo-pattern/app/model"
	"fiber-repo-pattern/app/repository"
	"fiber-repo-pattern/app/request"

	"github.com/mashingan/smapping"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) FindAll() []model.User {
	return service.userRepository.FindAll()
}

func (service *userService) FindById(UserId int) model.User {
	return service.userRepository.FindById(UserId)
}

func (service *userService) Create(dataUser request.UserCreateValidation) model.User {
	user := model.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&dataUser))
	if err != nil {
		panic("Failed Mapping Data")
	}
	result := service.userRepository.Create(user)
	return result
}

func (service *userService) Update(dataUser request.UserUpdateValidation) model.User {
	user := model.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&dataUser))
	if err != nil {
		panic("Failed Mapping Data")
	}
	result := service.userRepository.Update(user)
	return result
}

func (service *userService) Delete(dataUser model.User) error {
	return service.userRepository.Delete(dataUser)
}
