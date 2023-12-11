package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (service *UserRepository) Save(user *model.User) (*model.User, error) {
	result := service.db.Create(user)
	return user, result.Error
}

func (service *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := service.db.Find(&users)
	return users, result.Error
}

func (service *UserRepository) Find(id uuid.UUID) (*model.User, error) {
	var user *model.User
	result := service.db.Find(&user, "id = ?", id)

	if user.Id == uuid.Nil {
		return user, nil
	}

	return user, result.Error
}

func (service *UserRepository) FindByEmailAndPassword(username string, password string) (*model.User, error) {
	var user *model.User
	result := service.db.Find(&user, "email = ? and password = ?", username, password)

	if user.Id == uuid.Nil {
		return nil, errors.New("user not found")
	}

	return user, result.Error
}

func (service *UserRepository) Delete(id uuid.UUID) error {
	_, err := service.Find(id)
	if err != nil {
		return err
	}
	result := service.db.Delete(model.User{}, "id = ?", id)
	return result.Error
}
