package service

import (
	"github.com/google/uuid"
	"main.go/model"
	"main.go/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) FindAll() ([]model.User, error) {
	return service.userRepository.FindAll()
}

func (service *UserService) Find(id uuid.UUID) (*model.User, error) {
	return service.userRepository.Find(id)
}

func (service *UserService) FindByEmailAndPassword(email string, password string) (*model.User, error) {
	return service.userRepository.FindByEmailAndPassword(email, password)
}

func (service *UserService) Save(user *model.User) (*model.User, error) {
	return service.userRepository.Save(user)
}
