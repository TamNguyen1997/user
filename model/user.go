package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name             string     `json:"name" gorm:"type:string;not null"`
	Email            string     `json:"email" gorm:"type:string;not null"`
	Password         string     `json:"-" gorm:"type:string;not null"`
	Role             []UserRole `json:"-"`
	CreatedDate      time.Time  `json:"createdDate" gorm:"autoCreateTime"`
	LastModifiedDate time.Time  `json:"lastModifiedDate" gorm:"autoUpdateTime"`
}

type UserRole struct {
	Id               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Role             Role      `json:"name" gorm:"type:string;not null"`
	UserId           uuid.UUID `json:"userId" gorm:"type:uuid;not null"`
	CreatedDate      time.Time `json:"createdDate" gorm:"autoCreateTime"`
	LastModifiedDate time.Time `json:"lastModifiedDate" gorm:"autoUpdateTime"`
}
