package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"main.go/model"
)

type AuthorityRepository struct {
	db *gorm.DB
}

func NewAuthorityRepo(db *gorm.DB) *AuthorityRepository {
	return &AuthorityRepository{
		db: db,
	}
}

func (service *AuthorityRepository) Save(user *model.UserRole) (*model.UserRole, error) {
	result := service.db.Create(user)
	return user, result.Error
}

func (service *AuthorityRepository) Find(id uuid.UUID) (*model.UserRole, error) {
	var authority *model.UserRole
	result := service.db.Find(&authority, "id = ?", id)

	if authority.Id == uuid.Nil {
		return authority, nil
	}

	return authority, result.Error
}

func (service *AuthorityRepository) FindByEmailAndPassword(username string, password string) (*model.UserRole, error) {
	var authority *model.UserRole
	result := service.db.Find(&authority, "email = ? and password = ?", username, password)

	if authority.Id == uuid.Nil {
		return nil, errors.New("authority not found")
	}

	return authority, result.Error
}

func (service *AuthorityRepository) Delete(id uuid.UUID) error {
	_, err := service.Find(id)
	if err != nil {
		return err
	}
	result := service.db.Delete(model.UserRole{}, "id = ?", id)
	return result.Error
}
